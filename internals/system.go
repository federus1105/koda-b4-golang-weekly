package internals

import (
	"fmt"
	"sync"
	"time"
)

type OrderInterface interface {
	ShowMenu()
	AddOrder()
	ShowOrders()
	ProcessOrders()
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
	fmt.Printf("Pesanan %s, dengan jumlah %d berhasil. Total : %.2f\n", itemName, quantity, total)
}

func (o *OrderSystem) ShowOrders() {
	if len(o.Orders) == 0 {
		fmt.Println("Pesanan belum ada")
		return
	}
	fmt.Println("====   DAFTAR PESANAN  =====")
	var grandTotal float64
	for i, order := range o.Orders {
		fmt.Printf("%d. %s x %d - Rp %.2f\n", i+1, order.ItemName, order.Quantity, order.Total)
		grandTotal += order.Total
	}
	fmt.Printf("\nTotal Keseluruhan : Rp %.2f\n", grandTotal)
}

func (o *OrderSystem) ProcessOrders() {
	if len(o.Orders) == 0 {
		fmt.Println("Tidak ada pesanan yang diproses.")
		return
	}

	fmt.Println("Memproses Pesanan...")

	var wg sync.WaitGroup
	results := make(chan string, len(o.Orders))

	for _, ord := range o.Orders {
		wg.Add(1)
		go func(order Order) {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			results <- fmt.Sprintf("Pesanan %s, dengan jumlah %d selesai diproses!", order.ItemName, order.Quantity)
		}(ord)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("\n========================================")
		fmt.Println(res)
		fmt.Println("========================================")
	}

	fmt.Printf("\nSemua pesanan selesai!\n")
}
