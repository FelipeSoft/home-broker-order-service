package domain

type OrderRepository interface {
	GetOrdersByUserId(userId string) ([]*Order, error)
	GetOrderById(orderId string) (*Order, error)
	SaveOrder(order Order) (*int64, error)
}
