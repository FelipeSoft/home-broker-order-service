package domain

type OrderLimit struct {
	min_price float64
	max_price float64
	*Order
}

func NewOrderLimit(order *Order, min_price, max_price float64) *OrderLimit {
	return &OrderLimit{
		min_price: min_price,
		max_price: max_price,
	}
}
