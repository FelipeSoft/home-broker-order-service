package domain

import (
	"fmt"
	"log"

	"github.com/shopspring/decimal"
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
	ID       *int64
	UserId   string
	Ticker   string
	Quantity decimal.Decimal
	Price    decimal.Decimal
	MinPrice *decimal.Decimal
	MaxPrice *decimal.Decimal
	Total    decimal.Decimal
	Type     OrderType
	Status   OrderStatus
}

func NewOrder(ticker string, quantity, price string, minPrice, maxPrice *string, orderType OrderType, userId string) (*Order, error) {
	log.Printf("On Entity Order: %s", price)
	decimalQuantity, err := decimal.NewFromString(quantity)
	if err != nil {
		return nil, fmt.Errorf("decimal order quantity error: %s", err.Error())
	}

	decimalPrice, err := decimal.NewFromString(price)
	if err != nil {
		return nil, fmt.Errorf("decimal order price error: %s", err.Error())
	}

	order := &Order{
		Ticker:   ticker,
		Quantity: decimalQuantity,
		Price:    decimalPrice,
		Type:     orderType,
		Status:   OrderStatusPending,
		UserId:   userId,
	}
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