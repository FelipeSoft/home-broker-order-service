package mysql

import (
	"database/sql"

	"github.com/FelipeSoft/home-broker-order-service/internal/domain"
)

type OrderRepositoryMySql struct {
	db *sql.DB
}

func NewOrderRepositoryMySql(db *sql.DB) *OrderRepositoryMySql {
	return &OrderRepositoryMySql{
		db: db,
	}
}

func (r *OrderRepositoryMySql) GetOrdersByUserId(userId string) ([]interface{}, error) {
	return nil, nil
}

func (r *OrderRepositoryMySql) SaveOrder(order domain.Order) error {
// 	query := "INSERT INTO orders (user_id, ticker, quantity, ticker_price, total, type, order_type, min_price, max_price, status) VALUES (?,?,?,?,?,?,?,?,?,?);"
// 	r.db.Exec(query, 
// 	order.UserId, 
// 	order.Ticker,
// 	order.Quantity,
// 	order.Price,
// 	order.Total, 
// 	order.Type, 
// 	order
// )
	return nil
}

func (r *OrderRepositoryMySql) CancelOrder(orderId int64) error {
	return nil
}

func (r *OrderRepositoryMySql) FillOrder(orderId int64) error {
	return nil
}
