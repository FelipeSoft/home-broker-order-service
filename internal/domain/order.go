package domain

import (
	"fmt"
)

type OrderType string

const (
	OrderTypeBuy  OrderType = "BUY"
	OrderTypeSell OrderType = "SELL"
)

type OrderStatus string

const (
	OrderStatusPending  OrderStatus = "PENDING"
	OrderStatusFilled   OrderStatus = "FILLED"
	OrderStatusCanceled OrderStatus = "CANCELED"
)

type Order struct {
	ID       int64
	UserId   string
	Ticker   string
	Quantity float64
	Price    float64
	MinPrice *float64
	MaxPrice *float64
	Total    float64
	Type     OrderType
	Status   OrderStatus
}

func NewOrder(ticker string, quantity, price float64, minPrice, maxPrice *float64, orderType OrderType) (*Order, error) {
	if quantity <= 0 || price < 0 {
		return nil, fmt.Errorf("order ticker 'quantity' and 'price' must be positive values")
	}
	order := &Order{
		Ticker:   ticker,
		Quantity: quantity,
		Price:    price,
		Type:     orderType,
		Status:   OrderStatusPending,
	}
	order.calculateOrderTotal()
	return order, nil
}

func (o *Order) Fill() error {
	if o.Status == "CANCELED" {
		return fmt.Errorf("the order could not be filled because the current status is CANCELED")
	}
	if o.Status == "FILLED" {
		return fmt.Errorf("the order is already filled")
	}
	o.Status = OrderStatusPending
	return nil
}

func (o *Order) Cancel() error {
	if o.Status == "FILLED" {
		return fmt.Errorf("the order could not be CANCELED because the current status is FILLED")
	}
	if o.Status == "CANCELED" {
		return fmt.Errorf("the order is already CANCELED")
	}
	o.Status = OrderStatusCanceled
	return nil
}

func (o *Order) CheckAvailableBalance(balance float64) error {
	if o.Total < balance {
		return fmt.Errorf("insufficient balance for transaction")
	}
	return nil
}

func (o *Order) calculateOrderTotal() {
	o.Total = o.Price * o.Quantity
}
