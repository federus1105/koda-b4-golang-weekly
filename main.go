package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/federus1105/koda-b4-golang-weekly/internals"
	"github.com/joho/godotenv"
)

func main() {
	// ---  ENV LOAD ---
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed load .env", err)
	}
	// --- DEFAULT VALUE ---
	varTime := internals.DefaultEnv("TIME_CACHE", "15")
	times, err := strconv.Atoi(varTime)
	if err != nil {
		times = 15
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error:", r)
			main()
		}
	}()

	data, err := internals.GetData(time.Duration(times) * time.Second)
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
		fmt.Println("5. Bersihkan Cache")
		fmt.Println("\n0. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			newData, err := internals.GetData(15 * time.Second)
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
		case "5":
			fmt.Println("Membersihkan cache....")
			cachePath := filepath.Join(os.TempDir(), "data.json")
			if _, err := os.Stat(cachePath); err == nil {
				if err := os.Remove(cachePath); err != nil {
					fmt.Println("Gagal hapus:", err)
				} else {
					fmt.Println("Cache lama dihapus.")
				}
			} else {
				fmt.Println("File tidak ditemukan:", err)
			}
			newData, err := internals.GetData(15 * time.Second)
			if err != nil {
				fmt.Println("Gagal memperbarui data:", err)
			} else {
				system.Menu = newData
				fmt.Println("Menu berhasil diperbarui.")
			}
		case "0":
			fmt.Println("Terima kasih!")
			os.Exit(0)
		default:
			fmt.Printf("Pilihan tidak valid!\n")
		}
	}

}
