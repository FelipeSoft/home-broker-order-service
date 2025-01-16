package domain

type OrderMarket struct {
	MarketPrice float64
	*Order
}

type OrderMarketDTO struct {
	MarketPrice float64  `json:"market_price"`
	Order       OrderDTO `json:"order"`
}

func NewOrderMarket(order *Order, market_price float64) *OrderMarket {
	return &OrderMarket{
		MarketPrice: market_price,
		Order:       order,
	}
}
