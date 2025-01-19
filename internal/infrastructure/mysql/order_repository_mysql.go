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

func (r *OrderRepositoryMySql) GetOrdersByUserId(userId string) ([]*domain.Order, error) {
	query := "SELECT id, ticker, quantity, ticker_price, total, type, min_price, max_price, status FROM orders WHERE user_id = ;"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []*domain.Order
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(&order.ID, &order.MaxPrice, &order.MinPrice, &order.Price, &order.Quantity, &order.Status, &order.Ticker, &order.Total, &order.Type, &order.UserId)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (r *OrderRepositoryMySql) SaveOrder(order domain.Order) (*int64, error) {
	if order.ID == nil {
		query := "INSERT INTO orders (user_id, ticker, quantity, ticker_price, total, type, min_price, max_price, status) VALUES (?,?,?,?,?,?,?,?,?);"
		res, err := r.db.Exec(query,
			order.UserId,
			order.Ticker,
			order.Quantity,
			order.Price,
			order.Total,
			order.Type,
			order.MinPrice,
			order.MaxPrice,
			order.Status,
		)
		if err != nil {
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		order.ID = &id
		return order.ID, nil
	}
	query := "UPDATE orders SET user_id = ?, ticker = ?, quantity = ?, ticker_price = ?, total = ?, type = ?, min_price = ?, max_price = ?, status = ? WHERE id = ?;"
	_, err := r.db.Exec(query, order.UserId, order.Ticker, order.Quantity, order.Price, order.Total, order.Type, order.MinPrice, order.MaxPrice, order.Status, order.ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *OrderRepositoryMySql) GetOrderById(orderId string) (*domain.Order, error) {
	query := "SELECT id, user_id, ticker, quantity, ticker_price, total, type, min_price, max_price, status FROM orders WHERE id = ?;"
	rows, err := r.db.Query(query, orderId)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()
	var order domain.Order
	if rows.Next() {
		err = rows.Scan(&order.ID, &order.UserId, &order.Ticker, &order.Quantity, &order.Price, &order.Total, &order.Type, &order.MinPrice, &order.MaxPrice, &order.Status, &order.ID)
		if err != nil {
			return nil, err
		}
	}
	return &order, nil
}
