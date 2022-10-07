package database

import (
	"encoding/json"
	"os"
)

type Genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGensis(path string) (Genesis, error) {
	body, err := os.ReadFile(path)
	genesis := Genesis{Balances: make(map[Account]uint)}

	if err != nil {
		return genesis, err
	}

	var res map[string]interface{}
	json.Unmarshal(body, &res)

	balances := res["balances"].(map[string]interface{})

	for account, value := range balances {
		genesis.Balances = map[Account]uint{Account(account): uint(value.(float64))}
	}

	return genesis, err
}
