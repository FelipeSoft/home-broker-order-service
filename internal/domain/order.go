package domain

import (
	"fmt"
)

type Order struct {
	Ticker   string
	Quantity float64
	Price    float64
	total    float64
}

func NewOrder(ticker string, quantity, price float64) (*Order, error) {
	if quantity <= 0 || price <= 0 {
		return nil, fmt.Errorf("order ticker 'quantity' and 'price' must be positive values")
	}
	order := &Order{
		Ticker:   ticker,
		Quantity: quantity,
		Price:    price,
	}
	order.calculateOrderTotal()
	return order, nil
}

func (o *Order) CheckAvailableBalance(balance float64) error {
	if o.total < balance {
		return fmt.Errorf("insufficient balance for transaction")
	}
	return nil
}

func (o *Order) ChangeOrderQuantity(quantity float64) error {
	if quantity <= 0 {
		return fmt.Errorf("order ticker 'quantity' must be positive value")
	}
	o.Quantity = quantity
	o.calculateOrderTotal()
	return nil
}

func (o *Order) ChangeOrderPrice(price float64) error {
	if price <= 0 {
		return fmt.Errorf("order ticker 'price' must be positive value")
	}
	o.Price = price
	o.calculateOrderTotal()
	return nil
}

func (o *Order) calculateOrderTotal() {
	o.total = o.Price * o.Quantity
}
