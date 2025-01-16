package domain

type OrderLimit struct {
	MinPrice float64
	MaxPrice float64
	*Order
}

type OrderLimitDTO struct {
	MinPrice float64  `json:"min_price"`
	MaxPrice float64  `json:"max_price"`
	Order    OrderDTO `json:"order"`
}

func NewOrderLimit(order *Order, MinPrice, MaxPrice float64) *OrderLimit {
	return &OrderLimit{
		MinPrice: MinPrice,
		MaxPrice: MaxPrice,
		Order:    order,
	}
}
