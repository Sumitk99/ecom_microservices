package order

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"log"
)

type Repository interface {
	Close()
	PutOrder(ctx context.Context, o Order) error
	GetOrder(ctx context.Context, orderID, accountID string) (*Order, error)
	GetOrdersForAccount(ctx context.Context, accountID string) ([]*UserOrder, error)
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

func (r *postgresRepository) Close() {
	r.db.Close()
}

func TransactionHandler(tx *sql.Tx, err error) {
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
}

func (r *postgresRepository) PutOrder(ctx context.Context, order Order) error {
	var TransactionID any = nil
	if len(order.TransactionID) > 0 {
		TransactionID = order.TransactionID
	}

	var orderStatus string
	if order.PaymentStatus == "success" || order.PaymentStatus == "COD" {
		orderStatus = "Order Placed"
	} else if order.PaymentStatus == "pending" {
		orderStatus = "Payment Pending"
	} else {
		orderStatus = "Order Failed"
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer TransactionHandler(tx, err)
	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO orders(id, methodofpayment,transactionid, created_at, account_id, total_price, payment_status, order_status) VALUES ($1, $2, $3, $4,$5,$6,$7,$8)",
		order.ID,
		order.MethodOfPayment,
		TransactionID,
		order.CreatedAt,
		order.AccountID,
		order.TotalPrice,
		order.PaymentStatus,
		orderStatus,
	)
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, pq.CopyIn("ordered_products", "order_id", "product_id", "quantity", "name", "price", "image_url"))
	if err != nil {
		return err
	}

	for _, p := range order.Products {
		_, err = stmt.ExecContext(ctx, order.ID, p.ID, p.Quantity, p.Name, p.Price, p.ImageURL)
		if err != nil {
			return err
		}
	}
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil
}

func (r *postgresRepository) GetOrder(ctx context.Context, orderID, accountID string) (*Order, error) {
	order := new(Order)
	var transactionID sql.NullString

	orderDetails := r.db.QueryRowContext(
		ctx,
		"SELECT id, created_at, account_id, total_price, methodofpayment, transactionid, payment_status, order_status FROM orders WHERE id = $1 AND account_id = $2", orderID, accountID,
	)
	err := orderDetails.Scan(
		&order.ID,
		&order.CreatedAt,
		&order.AccountID,
		&order.TotalPrice,
		&order.MethodOfPayment,
		&transactionID,
		&order.PaymentStatus,
		&order.OrderStatus,
	)
	if transactionID.Valid {
		order.TransactionID = transactionID.String
	}
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while fetching order details : %s", err))
	}
	log.Printf("Order ID : %s\n", order.ID)
	log.Printf("Order Created At : %s\n", order.CreatedAt)
	log.Printf("Order Account ID : %s\n", order.AccountID)
	log.Printf("Order Total Price : %f\n", order.TotalPrice)
	log.Printf("Order Method Of Payment : %s\n", order.MethodOfPayment)
	log.Printf("Order Transaction ID : %s\n", order.TransactionID)
	log.Printf("Order Payment Status : %s\n", order.PaymentStatus)
	log.Printf("Order Order Status : %s\n", order.OrderStatus)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Unauthorized Access, Order Doesn't Belong to this Account or Order Doesn't Exists")
		}
	}

	orderedProducts, err := r.db.QueryContext(
		ctx,
		"SELECT product_id, quantity, name, price,image_url FROM ordered_products WHERE order_id = $1",
		orderID,
	)
	if orderedProducts == nil {
		return nil, errors.New("No Products Found for this order")
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for orderedProducts.Next() {
		product := new(OrderedProduct)

		if err := orderedProducts.Scan(&product.ID, &product.Quantity, &product.Name, &product.Price, &product.ImageURL); err != nil {
			log.Printf("Failed to scan ordered product: %v", err)
			continue
		}
		order.Products = append(order.Products, *product)
	}

	// Handle any errors returned by orderedProducts.Next()
	if err = orderedProducts.Err(); err != nil {
		log.Fatalf("Error occurred during iteration: %v", err)
		return nil, err
	}
	fmt.Printf("Products Number : %d\n", len(order.Products))
	for _, o := range order.Products {
		fmt.Println(o.ID, o.Quantity)
	}
	return order, nil
}

func (r *postgresRepository) GetOrdersForAccount(ctx context.Context, accountID string) ([]*UserOrder, error) {
	orders := []*UserOrder{}
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, created_at, total_price, order_status FROM orders WHERE account_id = $1", accountID,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		order := new(UserOrder)
		if err = rows.Scan(&order.OrderId, &order.CreatedAt, &order.TotalPrice, &order.OrderStatus); err != nil {
			log.Println(err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}
