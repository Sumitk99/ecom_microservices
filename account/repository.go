package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/account/models"
	_ "github.com/lib/pq"
	"log"
)

type Repository interface {
	Close() error
	SignUp(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
	ValidateNewAccount(ctx context.Context, email, phone string) (int, error)
	GetAccountByCredentials(ctx context.Context, email, phone string) (*Account, error)
	AddAddress(ctx context.Context, address *models.Address) error
	GetAddresses(ctx context.Context, userID string) ([]*models.Address, error)
	DeleteAddress(ctx context.Context, addressID, accountID string) error
	GetAddress(ctx context.Context, addressID, accountID string) (*models.Address, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() error {
	err := r.db.Close()
	return err
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) SignUp(ctx context.Context, a Account) error {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO accounts (
			id, name, password, email, phone, token, user_type, refresh_token, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		a.ID, a.Name, a.Password, a.Email, a.Phone, a.Token, a.UserType, a.RefreshToken, a.CreatedAt, a.UpdatedAt,
	)
	log.Println(fmt.Sprintf("Account inserted: ID=%s, Name=%s, Email=%s\n Error=%v", a.ID, a.Name, a.Email, err))
	return err
}

func (r *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	log.Println("Repo Side: ", id)

	row := r.db.QueryRowContext(ctx, "SELECT id, name, email, phone, user_type FROM accounts WHERE id = $1", id)
	a := &Account{}
	if err := row.Scan(&a.ID, &a.Name, &a.Email, &a.Phone, &a.UserType); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no entries of any account found") // No rows found
		}
		return nil, err
	}
	return a, nil
}

func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, nam FROM accounts ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	accounts := []Account{}
	for rows.Next() {
		a := Account{}
		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *postgresRepository) ValidateNewAccount(ctx context.Context, email, phone string) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM accounts WHERE email = $1 OR phone = $2", email, phone).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *postgresRepository) GetAccountByCredentials(ctx context.Context, Email, Phone string) (*Account, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, password, email,phone, user_type FROM accounts WHERE email = $1 OR phone = $2", Email, Phone)
	acc := Account{}
	if err := row.Scan(
		&acc.ID,
		&acc.Name,
		&acc.Password,
		&acc.Email,
		&acc.Phone,
		&acc.UserType,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No Account Found For Given Email or Phone")
		}
		return nil, err
	}
	return &acc, nil
}

func (r *postgresRepository) AddAddress(ctx context.Context, address *models.Address) error {
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM accounts WHERE id = $1)", address.UserID).Scan(&exists)
	if err != nil || !exists {
		address.IsDefault = true
	} else {
		address.IsDefault = false
	}
	query := `
		INSERT INTO addresses (
			id, account_id, is_default, street_address, apartment_unit, 
			city, state, country, zip_code, created_at, name, phone
		) VALUES (
			$1, $2, $3, $4, $5, 
			$6, $7, $8, $9, $10, $11, $12
		);
	`

	_, err = r.db.ExecContext(ctx, query,
		address.AddressID,
		address.UserID,
		address.IsDefault,
		address.Street,
		address.ApartmentUnit,
		address.City,
		address.State,
		address.Country,
		address.ZipCode,
		address.CreatedAt,
		address.Name,
		address.Phone,
	)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (r *postgresRepository) GetAddresses(ctx context.Context, userID string) ([]*models.Address, error) {
	query := `
		SELECT id, account_id, is_default, street_address, apartment_unit, 
			city, state, country, zip_code, created_at, name, phone
		FROM addresses
		WHERE account_id = $1;
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	if errors.Is(rows.Err(), sql.ErrNoRows) {
		return nil, errors.New("No addresses found for this user")
	}
	defer rows.Close()

	addresses := []*models.Address{}
	for rows.Next() {
		address := &models.Address{}
		err = rows.Scan(
			&address.AddressID,
			&address.UserID,
			&address.IsDefault,
			&address.Street,
			&address.ApartmentUnit,
			&address.City,
			&address.State,
			&address.Country,
			&address.ZipCode,
			&address.CreatedAt,
			&address.Name,
			&address.Phone,
		)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (r *postgresRepository) DeleteAddress(ctx context.Context, addressID, accountID string) error {
	query := `
		DELETE FROM addresses
		WHERE id = $1 AND account_id = $2
	`
	res, err := r.db.ExecContext(ctx, query, addressID, accountID)
	if err != nil {
		return errors.New(err.Error())
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Invalid address Id for this user")
	}
	return nil
}

func (r *postgresRepository) GetAddress(ctx context.Context, addressID, accountID string) (*models.Address, error) {
	query := `
		SELECT id, account_id, is_default, street_address, apartment_unit, 
			city, state, country, zip_code, created_at, name, phone
		FROM addresses
		WHERE id = $1 AND account_id = $2;
	`

	row := r.db.QueryRowContext(ctx, query, addressID, accountID)
	address := &models.Address{}
	err := row.Scan(
		&address.AddressID,
		&address.UserID,
		&address.IsDefault,
		&address.Street,
		&address.ApartmentUnit,
		&address.City,
		&address.State,
		&address.Country,
		&address.ZipCode,
		&address.CreatedAt,
		&address.Name,
		&address.Phone,
	)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return address, nil
}
