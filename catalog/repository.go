package catalog

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

var (
	ErrNotFound = errors.New("Product not found")
)

const indexName = "catalog"

type Repository interface {
	Close() error
	PutProduct(ctx context.Context, p Product) error
	GetProductByID(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	ListProductWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type elasticRepository struct {
	client *elasticsearch.Client
}

type productDocument struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func NewElasticRepository(cloudId, apiKey string) (Repository, error) {

	client, err := elasticsearch.NewClient(
		elasticsearch.Config{
			CloudID: cloudId,
			APIKey:  apiKey,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error creating elasticsearch client: %w", err)
	}
	_, err = client.Info()
	if err != nil {
		return nil, fmt.Errorf("error getting cluster info: %w", err)
	}
	return &elasticRepository{client: client}, nil
}

func (r *elasticRepository) Close() error {
	return nil
}

func (r *elasticRepository) PutProduct(ctx context.Context, p Product) error {
	doc := productDocument{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	jsonDoc, err := json.Marshal(doc)
	if err != nil {
		log.Fatalf("Error marshalling document: %s", err)
	}
	res, err := r.client.Index(
		indexName,
		strings.NewReader(string(jsonDoc)),
		r.client.Index.WithDocumentID(p.ID),
		r.client.Index.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error indexing document: %s", res.String())
	}
	return err
}

func (r *elasticRepository) GetProductByID(ctx context.Context, id string) (*Product, error) {
	res, err := r.client.Get(
		"catalog",
		id,
		r.client.Get.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 404 {
		return nil, ErrNotFound
	}

	//Raw Response → Generic Map → Source Map → JSON Bytes → Product Struct
	var body map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	source, ok := body["_source"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error accessing _source in response")
	}

	p := productDocument{}
	sourceBytes, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("error marshaling _source: %w", err)
	}

	if err := json.Unmarshal(sourceBytes, &p); err != nil {
		return nil, fmt.Errorf("error unmarshaling into product: %w", err)
	}

	return &Product{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, nil
}

func (r *elasticRepository) ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	// Create the search query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"from": skip,
		"size": take,
	}

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error marshaling query: %w", err)
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("catalog"),
		r.client.Search.WithBody(strings.NewReader(string(jsonQuery))),
	)
	if err != nil {
		return nil, fmt.Errorf("error performing search: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error searching products: %s", res.String())
	}

	var searchResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	hits, ok := searchResult["hits"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error accessing hits in response")
	}

	hitsList, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error accessing hits list in response")
	}

	products := make([]Product, 0, len(hitsList))
	for _, hit := range hitsList {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error converting hit to map")
		}

		id, _ := hitMap["_id"].(string)
		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error accessing _source in hit")
		}

		sourceBytes, err := json.Marshal(source)
		if err != nil {
			return nil, fmt.Errorf("error marshaling source: %w", err)
		}

		var p productDocument
		if err := json.Unmarshal(sourceBytes, &p); err != nil {
			return nil, fmt.Errorf("error unmarshaling product: %w", err)
		}

		products = append(products, Product{
			ID:          id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return products, nil
}

func (r *elasticRepository) ListProductWithIDs(ctx context.Context, ids []string) ([]Product, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"ids": map[string]interface{}{
				"values": ids,
			},
		},
	}

	searchBody, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error marshaling query: %w", err)
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("catalog"),
		r.client.Search.WithBody(bytes.NewReader(searchBody)),
		r.client.Search.WithSize(len(ids)),
	)
	if err != nil {
		return nil, fmt.Errorf("error searching documents: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search response error: %s", res.String())
	}

	var searchRes map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchRes); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}
	hits, ok := searchRes["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error parsing hits from response")
	}
	products := make([]Product, 0, len(hits))
	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})

		product := Product{
			ID:          hitMap["_id"].(string),
			Name:        source["name"].(string),
			Description: source["description"].(string),
			Price:       source["price"].(float64),
		}
		products = append(products, product)
	}

	return products, nil
}
func (r *elasticRepository) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	// Create the search query
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"name", "description"},
			},
		},
	}

	searchBody, err := json.Marshal(searchQuery)
	if err != nil {
		return nil, fmt.Errorf("error marshaling query: %w", err)
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("catalog"),
		r.client.Search.WithBody(bytes.NewReader(searchBody)),
		r.client.Search.WithFrom(int(skip)),
		r.client.Search.WithSize(int(take)),
	)
	if err != nil {
		return nil, fmt.Errorf("error searching products: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search response error: %s", res.String())
	}

	var searchRes map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchRes); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	hits, ok := searchRes["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error parsing hits from response")
	}

	products := make([]Product, 0, len(hits))
	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})

		product := Product{
			ID:          hitMap["_id"].(string),
			Name:        source["name"].(string),
			Description: source["description"].(string),
			Price:       source["price"].(float64),
		}
		products = append(products, product)
	}

	return products, nil
}
