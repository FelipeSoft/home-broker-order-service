package domain

type MarketOrder struct {
	*Order
}

func NewMarketOrder(order *Order) *MarketOrder {
	return &MarketOrder{
		Order:       order,
	}
}
