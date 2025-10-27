package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// === FETCHING DATA ====
func LoadDataFromSource() ([]MenuItem, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/federus1105/koda-b4-golang-weekly-data/refs/heads/main/data.json")
	if err != nil {
		fmt.Println("Fecth data failed!")
	}

	var menu []MenuItem
	body, err := io.ReadAll(
		resp.Body,
	)

	if err != nil {
		fmt.Println("Failed to read body")
	}

	json.Unmarshal(body, &menu)
	system := &OrderSystem{
		Menu: menu,
	}
	_ = system
	return menu, nil
}
