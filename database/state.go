package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type State struct {
	Balances  map[Account]uint
	txMempool []Tx

	dbFile *os.File
}

func NewStateFromDisk(genFilePath string, txDbFilePath string) (*State, error) {
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	return nil, err
	// }

	// genFilePath := filepath.Join(cwd, "genesis.json")
	genFile, err := loadGensis(genFilePath)
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range genFile.Balances {
		balances[account] = balance
	}

	// txDbFilePath := filepath.Join(cwd, "tx.db")
	f, err := os.OpenFile(txDbFilePath, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	state := &State{Balances: balances, txMempool: make([]Tx, 0), dbFile: f}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var tx Tx
		json.Unmarshal(scanner.Bytes(), &tx)

		if err := state.apply(tx); err != nil {
			return nil, err
		}
	}

	return state, err
}

func (s *State) apply(tx Tx) error {
	if tx.isReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	if tx.Value > s.Balances[tx.From] {
		return fmt.Errorf("insufficient balance")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value

	return nil
}
