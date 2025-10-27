package internals

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
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
