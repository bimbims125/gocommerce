package repository

import (
	"context"
	"database/sql"
	"gocommerce/internal/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *entity.Order) (int, error) {
	query := "INSERT INTO orders (user_id, product_id, quantity, total_price, transaction_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, order.UserID, order.ProductID, order.Quantity, order.TotalPrice, order.TransactionID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]entity.GetOrder, error) {
	query := `SELECT
							o.id,
							o.user_id,
							o.product_id,
							o.quantity,
							o.total_price,
							o.transaction_id,
							u.name AS user_name,
							u.email AS user_email,
							p.name AS product_name,
							p.image_url AS image_url
						FROM
							orders o
						JOIN
							users u ON o.user_id = u.id
						JOIN
							products p ON o.product_id = p.id;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []entity.GetOrder
	for rows.Next() {
		var order entity.GetOrder
		if err := rows.Scan(&order.ID, &order.User.ID,
			&order.Product.ID, &order.Quantity, &order.TotalPrice,
			&order.TransactionID, &order.User.Name, &order.User.Email, &order.Product.Name, &order.Product.ImageURL); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) UpdateOrderPaymentStatus(ctx context.Context, transactionID string, paymentStatus string) error {
	query := "UPDATE orders SET payment_status = $1 WHERE transaction_id = $2"
	_, err := r.db.ExecContext(ctx, query, paymentStatus, transactionID)
	if err != nil {
		return err
	}
	return nil
}
