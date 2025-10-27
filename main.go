package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/federus1105/koda-b4-golang-weekly/internals"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error:", r)
			main()
		}
	}()

	data, err := internals.GetData(15 * time.Millisecond)
	if err != nil {
		fmt.Println(err)
	}
	system := &internals.OrderSystem{
		Menu: data,
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n===== SISTEM PEMESANAN =====\n")
		fmt.Println("1. Lihat Menu")
		fmt.Println("2. Tambah Pesanan")
		fmt.Println("3. Lihat Pesanan")
		fmt.Println("4. Proses Pesanan")
		fmt.Println("\n0. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			system.ShowMenu()
		case "2":
			system.ShowMenu()
			fmt.Printf("Masukkan nama item: ")
			item, _ := reader.ReadString('\n')
			item = strings.TrimSpace(item)

			fmt.Printf("Masukkan jumlah: ")
			qtyInput, _ := reader.ReadString('\n')
			qtyInput = strings.TrimSpace(qtyInput)
			qty, err := strconv.Atoi(qtyInput)
			if err != nil {
				panic("Jumlah tidak valid")
			}
			system.AddOrder(item, qty)
		case "3":
			system.ShowOrders()
		case "4":
			system.ProcessOrders()
		case "0":
			fmt.Println("Terima kasih!")
			os.Exit(0)
		default:
			fmt.Printf("Pilihan tidak valid!\n")
		}
	}

}
