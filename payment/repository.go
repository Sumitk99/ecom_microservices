package payment

import (
	"database/sql"
	"github.com/Sumitk99/ecom_microservices/payment/models"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Repository interface {
	Close() error
	ListCards(userID string) ([]*models.Card, error)
	AddCard(card *models.Card) error
	RemoveCard(cardID string) error
	GetCard(cardID string) (*models.Card, error)
	SaveTransaction(userID, cardID string, amount float64, currency, description string, result *models.PaymentResult) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &postgresRepository{db: db}, nil
}

func (r *postgresRepository) Close() error {
	r.db.Close()
	return nil
}

func (r *postgresRepository) AddCard(card *models.Card) error {
	query := `
		INSERT INTO cards (id, user_id, card_number, card_holder_name, expiry_month, expiry_year, card_type, is_default, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Exec(query,
		card.ID,
		card.UserID,
		card.CardNumber,
		card.CardHolderName,
		card.ExpiryMonth,
		card.ExpiryYear,
		card.CardType,
		card.IsDefault,
		time.Now(),
	)
	return err
}

func (r *postgresRepository) RemoveCard(cardID string) error {
	query := `DELETE FROM cards WHERE id = $1`
	_, err := r.db.Exec(query, cardID)
	return err
}

func (r *postgresRepository) GetCard(cardID string) (*models.Card, error) {
	query := `SELECT id, user_id, card_number, card_holder_name, expiry_month, expiry_year, card_type, is_default, created_at FROM cards WHERE id = $1`
	row := r.db.QueryRow(query, cardID)

	card := &models.Card{}
	err := row.Scan(
		&card.ID,
		&card.UserID,
		&card.CardNumber,
		&card.CardHolderName,
		&card.ExpiryMonth,
		&card.ExpiryYear,
		&card.CardType,
		&card.IsDefault,
		&card.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (r *postgresRepository) ListCards(userID string) ([]*models.Card, error) {
	query := `SELECT id, user_id, card_number, card_holder_name, expiry_month, expiry_year, card_type, is_default, created_at FROM cards WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []*models.Card
	for rows.Next() {
		card := &models.Card{}
		err := rows.Scan(
			&card.ID,
			&card.UserID,
			&card.CardNumber,
			&card.CardHolderName,
			&card.ExpiryMonth,
			&card.ExpiryYear,
			&card.CardType,
			&card.IsDefault,
			&card.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *postgresRepository) SaveTransaction(userID, cardID string, amount float64, currency, description string, result *models.PaymentResult) error {
	query := `
		INSERT INTO transactions (
			transaction_id, user_id, card_id, amount, currency, description, status, timestamp
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(
		query,
		result.TransactionID,
		userID,
		cardID,
		amount,
		currency,
		description,
		result.Status,
		result.Timestamp,
	)

	return err
}
