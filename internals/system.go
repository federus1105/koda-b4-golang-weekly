package internals

import (
	"fmt"
)

type OrderInterface interface {
	ShowMenu()
	AddOrder()
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

func (o *OrderSystem) AddOrder(itemName string, quantity int) {
	var selected *MenuItem
	for _, item := range o.Menu {
		if item.Name == itemName {
			selected = &item
			break
		}
	}
	if selected == nil {
		panic("Masukkan nama item dengan benar!")
	}
	total := selected.Price * float64(quantity)
	o.Orders = append(o.Orders, Order{
		ItemName: itemName,
		Quantity: quantity,
		Total:    total,
	})
	fmt.Printf("Pesanan %s, dengan jumlah %d berhasil. Total : %f\n", itemName, quantity, total)
}
