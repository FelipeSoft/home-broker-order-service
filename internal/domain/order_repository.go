package domain

type OrderRepository interface {
	GetOrdersByUserId(userId string) ([]Order, error)
	SaveOrder(order Order) error
	ChangeTotal(order Order) error
	Cancel(orderId int64) error
	Fill(orderId int64) error
}
