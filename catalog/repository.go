package catalog

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/catalog/models"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
	"io"

	"log"
	"strings"
)

var (
	ErrNotFound = errors.New("Product not found")
)

const indexName = "catalog"

type Repository interface {
	Close() error
	PutProduct(ctx context.Context, p models.Product) error
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]models.Product, error)
	ListProductWithIDs(ctx context.Context, ids []string) ([]models.Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]models.Product, error)
}

type OpenSearchRepository struct {
	client *opensearch.Client
}

func NewOpenSearchRepository(AwsEndPoint, AwsRegion string) (Repository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AwsRegion))
	if err != nil {
		log.Fatalf("Error loading AWS config: %v", err)
	}

	// Create OpenSearch Client with AWS SigV4 signer
	signer, err := awsv2.NewSigner(cfg)
	if err != nil {
		log.Fatalf("Error creating signer: %v", err)
	}

	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{AwsEndPoint},
		Signer:    signer, // Use AWS Signature V4
	})

	if err != nil {
		return nil, err
	}
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting cluster info: %v", err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println("Cluster Info:", string(body))
	return &OpenSearchRepository{
		client: client,
	}, nil
}

func (r *OpenSearchRepository) Close() error {
	return nil
}

func (r *OpenSearchRepository) PutProduct(ctx context.Context, p models.Product) error {
	doc := models.ProductDocument{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		ImageUrl:    p.ImageUrl,
		SellerID:    p.SellerID,
		SellerName:  p.SellerName,
		Category:    p.Category,
		Stock:       p.Stock,
		Locations:   p.Locations,
		Sizes:       p.Sizes,
		Colors:      p.Colors,
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

func (r *OpenSearchRepository) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	res, err := r.client.Get(
		indexName,
		id,
		r.client.Get.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, ErrNotFound
	}

	var body map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	source, ok := body["_source"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error accessing _source in response")
	}

	product := models.ProductDocument{}
	sourceBytes, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("error marshaling _source: %w", err)
	}

	if err := json.Unmarshal(sourceBytes, &product); err != nil {
		return nil, fmt.Errorf("error unmarshaling into product: %w", err)
	}

	return &models.Product{
		ID:          id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl,
		SellerID:    product.SellerID,
		SellerName:  product.SellerName,
		Category:    product.Category,
		Stock:       product.Stock,
		Locations:   product.Locations,
		Sizes:       product.Sizes,
		Colors:      product.Colors,
	}, nil
}

func (r *OpenSearchRepository) ListProducts(ctx context.Context, skip uint64, take uint64) ([]models.Product, error) {
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
		r.client.Search.WithIndex(indexName),
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

	products := make([]models.Product, 0, len(hitsList))
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

		var p models.ProductDocument
		if err := json.Unmarshal(sourceBytes, &p); err != nil {
			return nil, fmt.Errorf("error unmarshaling product: %w", err)
		}

		products = append(products, models.Product{
			ID:         id,
			Name:       p.Name,
			Price:      p.Price,
			SellerName: p.SellerName,
			ImageUrl:   p.ImageUrl,
		})
	}

	return products, nil
}

func (r *OpenSearchRepository) ListProductWithIDs(ctx context.Context, ids []string) ([]models.Product, error) {
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
		r.client.Search.WithIndex(indexName),
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

	products := make([]models.Product, 0, len(hits))
	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})

		product := models.Product{
			ID:         hitMap["_id"].(string),
			Name:       source["name"].(string),
			Price:      source["price"].(float64),
			ImageUrl:   source["image_url"].(string),
			SellerName: source["seller_name"].(string),
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *OpenSearchRepository) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]models.Product, error) {
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
		r.client.Search.WithIndex(indexName),
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

	products := make([]models.Product, 0, len(hits))
	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})

		product := models.Product{
			ID:         hitMap["_id"].(string),
			Name:       source["name"].(string),
			Price:      source["price"].(float64),
			ImageUrl:   source["image_url"].(string),
			SellerName: source["seller_name"].(string),
		}
		products = append(products, product)
	}

	return products, nil
}
