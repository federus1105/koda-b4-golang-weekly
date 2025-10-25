package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/federus1105/koda-b4-golang-weekly/internals"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error:", r)
			main()
		}
	}()

	system := &internals.OrderSystem{
		Menu: []internals.MenuItem{
			{Name: "Ayam Geprek", Price: 22000, Type: "Makanan"},
			{Name: "Ayam Geprek", Price: 22000, Type: "Makanan"},
			{Name: "Ayam Geprek", Price: 22000, Type: "Makanan"},
		},
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("===== SISTEM PEMESANAN =====\n")
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
		case "0":
			fmt.Println("Terima kasih!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid!\n")
		}
	}

}
