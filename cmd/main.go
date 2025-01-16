package main

import (
	"fmt"

	"github.com/FelipeSoft/home-broker-order-service/internal/domain"
)

func main() {
	order, err := domain.NewOrder("PETR4", 3, 30.20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(order);
}