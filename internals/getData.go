package internals

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func GetData(cacheDuration time.Duration) ([]MenuItem, error) {
	// === JALUR FILE SEMENTARA ===
	cacheDir := os.TempDir()
	cacheFile := filepath.Join(cacheDir, "data.json")

	// === PERIKSA FILE CACHE ADA DAN MASIH VALID ===
	fmt.Println("Loadingg.......")
	fileInfo, err := os.Stat(cacheFile)
	if err == nil && time.Since(fileInfo.ModTime()) < cacheDuration {
		// === CACHE VALID BACA DARI FILE ===
		fmt.Printf("Cache valid. Mengambil data dari cache (%v lalu)\n", time.Since(fileInfo.ModTime()).Round(time.Second))
		data, err := os.ReadFile(cacheFile)
		if err != nil {
			return nil, err
		}
		var menu []MenuItem
		err = json.Unmarshal(data, &menu)
		return menu, err
	}

	//  === JALANKAN JIKA CACHE TIDAK VALID ===
	menu, err := LoadDataFromSource()
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(menu)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(cacheFile, jsonData, 0644)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func InitDB() ([]MenuItem, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed load .env", err)
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Failed to connect", err)
	}

	var Menu MenuItem

	err = conn.QueryRow(context.Background(),
		"SELECT id, name, price, type FROM product",
	).Scan(&Menu.Id, &Menu.Name, &Menu.Price, &Menu.Type)

	if err != nil {
		fmt.Println("Error map database")
	}

	rows, err := conn.Query(context.Background(),
		"SELECT id, name, price, type FROM product",
	)
	if err != nil {
		fmt.Println("failed get database")
	}

	product, err := pgx.CollectRows(rows, pgx.RowToStructByName[MenuItem])
	if err != nil {
		fmt.Println("failed to map")
	}
	return product, nil
}
