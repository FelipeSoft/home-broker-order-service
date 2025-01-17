package domain

type LimitOrder struct {
	MinPrice float64
	MaxPrice float64
	*Order
}

func NewLimitOrder(order *Order, MinPrice, MaxPrice float64) *LimitOrder {
	return &LimitOrder{
		MinPrice: MinPrice,
		MaxPrice: MaxPrice,
		Order:    order,
	}
}
