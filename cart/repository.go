package cart

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/cart/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

const NoUserData = "not enough user Data provided"

type Repository interface {
	Close()
	AddItem(ctx context.Context, cartName, accountId, guestId, productId string, quantity uint64) error
	DeleteItem(ctx context.Context, cartName, accountId, guestId, productId string) error
	GetCartItems(ctx context.Context, cartName, accountId, guestId string) ([]models.CartItem, error)
	UpdateItem(ctx context.Context, cartName, accountId, guestId, productId string, updatedQuantity uint64) error
	DeleteCart(ctx context.Context, cartName, accountId, guestId string) error
	//MergeGuestCart(ctx context.Context, cartName,accountId,  guestId string) error
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) Close() {
	r.db.Close()
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

func (r *postgresRepository) AddItem(ctx context.Context, cartName, accountId, guestId, productId string, quantity uint64) error {
	var err error
	fmt.Printf("guestId : %s\n", guestId)
	if len(accountId) > 0 && len(cartName) > 0 {
		_, err = r.db.ExecContext(ctx, `
		INSERT INTO cart (
			cartName, accountId, productId, quantity
		) VALUES ($1, $2, $3, $4)`,
			cartName, accountId, productId, quantity,
		)
	} else if len(guestId) > 0 {
		fmt.Printf("guestId : %s\n", guestId)
		_, err = r.db.ExecContext(ctx, `
		INSERT INTO guestCart (
		    guestId, productId, quantity
		) VALUES ($1, $2, $3)`,
			guestId, productId, quantity,
		)
	} else {
		return errors.New(NoUserData)
	}
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return errors.New("Item already exists in selected cart")
			}
		}
	}

	return err
}

func (r *postgresRepository) GetCartItems(ctx context.Context, cartName, accountId, guestId string) ([]models.CartItem, error) {
	if accountId == "" && guestId == "" {
		log.Println("No info provided")
		return nil, errors.New("no info provided")
	}
	var err error
	var rows *sql.Rows

	if len(accountId) > 0 && len(cartName) > 0 {
		rows, err = r.db.QueryContext(ctx, `
			SELECT productId, quantity FROM cart WHERE accountId = $1 AND cartName = $2`, accountId, cartName)
	} else if len(guestId) > 0 {
		rows, err = r.db.QueryContext(ctx, `
			SELECT productId, quantity FROM guestCart WHERE guestId = $1`, guestId)
	} else {
		return nil, errors.New(NoUserData)
	}

	defer rows.Close()
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("error getting cartItems : %s", err))
	}

	cartProducts := []models.CartItem{}
	for rows.Next() {
		a := models.CartItem{}
		if err := rows.Scan(&a.ProductID, &a.Quantity); err != nil {
			return nil, err
		}
		cartProducts = append(cartProducts, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cartProducts, err
}

func (r *postgresRepository) DeleteItem(ctx context.Context, cartName, accountId, guestId, productId string) error {
	var err error
	if len(accountId) > 0 && len(cartName) > 0 {
		_, err = r.db.ExecContext(ctx, `DELETE FROM cart WHERE cartName = $1 AND productId = $2 AND accountId = $3`, cartName, productId, accountId)
	} else if len(guestId) > 0 {
		_, err = r.db.ExecContext(ctx, `DELETE FROM guestcart WHERE productId = $1 AND guestId = $2`, productId, guestId)
	} else {
		return errors.New(NoUserData)
	}
	if err != nil {
		log.Println(err)
		return errors.New(fmt.Sprintf("error deleting item from cart : %s", err))
	}
	return nil
}

func (r *postgresRepository) UpdateItem(ctx context.Context, cartName, accountId, guestId, productId string, updatedQuantity uint64) error {
	var err error
	if len(accountId) > 0 && len(cartName) > 0 {
		_, err = r.db.ExecContext(ctx, `UPDATE cart SET quantity = $1 WHERE cartName = $2 AND productId = $3 AND accountId = $4`, updatedQuantity, cartName, productId, accountId)
	} else if len(guestId) > 0 {
		_, err = r.db.ExecContext(ctx, `UPDATE guestcart SET quantity = $1 WHERE guestId = $2 AND productId = $3`, updatedQuantity, guestId, productId)
	} else {
		return errors.New(NoUserData)
	}
	if err != nil {
		log.Println(err)
		return errors.New(fmt.Sprintf("error updating item in cart : %s", err))
	}
	return nil
}

func (r *postgresRepository) DeleteCart(ctx context.Context, cartName, accountId, guestId string) error {
	var err error
	if len(accountId) > 0 && len(cartName) > 0 {
		_, err = r.db.ExecContext(ctx, `DELETE FROM cart WHERE cartName = $1 AND accountId = $2`, cartName, accountId)
	} else if len(guestId) > 0 {
		_, err = r.db.ExecContext(ctx, `DELETE FROM guestcart WHERE guestId = $1`, guestId)
	} else {
		return errors.New(NoUserData)
	}
	if err != nil {
		log.Println(err)
		return errors.New(fmt.Sprintf("error deleting cart : %s", err))
	}

	return nil
}
