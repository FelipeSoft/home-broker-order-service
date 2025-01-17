package domain

import "fmt"

type Order struct {
	Ticker         string
	Quantity       float64
	EstimatedPrice float64
	total          float64
}

func NewOrder(ticker string, quantity, estimated_price float64) (*Order, error) {
	if quantity <= 0 || estimated_price <= 0 {
		return nil, fmt.Errorf("order ticker 'quantity' and 'estimated_price' must be positive values")
	}
	order := &Order{
		Ticker:         ticker,
		Quantity:       quantity,
		EstimatedPrice: estimated_price,
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

func (o *Order) ChangeOrderEstimatedPrice(estimated_price float64) error {
	if estimated_price <= 0 {
		return fmt.Errorf("order ticker 'estimated_price' must be positive value")
	}
	o.EstimatedPrice = estimated_price
	o.calculateOrderTotal()
	return nil
}

func (o *Order) calculateOrderTotal() {
	o.total = o.EstimatedPrice * o.Quantity
}
