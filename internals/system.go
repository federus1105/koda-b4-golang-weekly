package internals

import (
	"fmt"
)

type OrderInterface interface {
	ShowMenu()
}

type Order struct {
	ItemName string
	Quantity int
	Total    float64
}

type OrderSystem struct {
	Menu   []MenuItem
	Orders []Order
}

func (o *OrderSystem) ShowMenu() {
	fmt.Println("=====   DAFTAR MENU   ======")
	for i, item := range o.Menu {
		fmt.Printf("%d. %s (%s) - Rp %.2f\n", i+1, item.Name, item.Type, item.Price)
	}
	fmt.Println()
}
