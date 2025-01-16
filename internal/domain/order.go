package domain

import "fmt"

type Order struct {
	ticker   string
	quantity float64
	price    float64
	total    float64
	// status -> open | in process | canceled | shipped complete
}

func NewOrder(ticker string, quantity, price float64) (*Order, error) {
	if quantity <= 0 || price <= 0 {
		return nil, fmt.Errorf("order quantity and price must be positive values")
	}
	order := &Order{
		ticker:   ticker,
		quantity: quantity,
		price:    price,
	}
	order.calculateOrderTotal()
	return order, nil
}

func (o *Order) ExecuteOrder() error {
	return nil
}

func (o *Order) ChangeOrderQuantity(quantity float64) error {
	if quantity <= 0 {
		return fmt.Errorf("order quantity must be positive value")
	}
	o.quantity = quantity
	return nil
}

func (o *Order) ChangeOrderPrice(price float64) error {
	if price <= 0 {
		return fmt.Errorf("order price must be positive value")
	}
	o.price = price
	return nil
}

func (o *Order) calculateOrderTotal() float64 {
	total := o.price * o.quantity
	return total
}

func (o *Order) Ticker() string {
	return o.ticker
}

func (o *Order) Price() float64 {
	return o.price
}

func (o *Order) Quantity() float64 {
	return o.quantity
}

func (o *Order) Total() float64 {
	return o.total
}
