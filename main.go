package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/federus1105/koda-b4-golang-weekly/internals"
)

func main() {
	product, err := internals.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	system := &internals.OrderSystem{
		Menu: product,
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
			newData, err := internals.InitDB()
			if err != nil {
				fmt.Println("Gagal memuat data:", err)
				break
			}
			system.Menu = newData
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
