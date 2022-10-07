package database

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewStateFromDisk(t *testing.T) {
	cwd, _ := os.Getwd()

	t.Run("proper genesis and transactions files", func(t *testing.T) {
		genFilePath := filepath.Join(cwd, "test_files", "genesis_statefull.json")
		txDbFilePath := filepath.Join(cwd, "test_files", "tx.db")

		_, err := NewStateFromDisk(genFilePath, txDbFilePath)

		assertNoError(t, err)
	})
}

func TestApply(t *testing.T) {
	t.Run("enough balance to apply the transcation", func(t *testing.T) {
		state := State{Balances: map[Account]uint{"nick": 1_000_000, "alina": 100_000}}

		err := state.apply(Tx{From: "nick", To: "alina", Value: 100_000})

		assertNoError(t, err)

		want := State{Balances: map[Account]uint{"nick": 900_000, "alina": 200_000}}

		assertBalances(t, want, state)
	})

	t.Run("not enough balance to apply the transaction", func(t *testing.T) {
		state := State{Balances: map[Account]uint{"nick": 0, "alina": 0}}

		err := state.apply(Tx{From: "nick", To: "alina", Value: 100})

		assertError(t, err)

		want := State{Balances: map[Account]uint{"nick": 0, "alina": 0}}

		assertBalances(t, want, state)
	})

	t.Run("reward transaction", func(t *testing.T) {
		state := State{Balances: map[Account]uint{"nick": 100_000, "alina": 0}}

		err := state.apply(Tx{From: "nick", To: "alina", Value: 1_000_000, Data: "reward"})

		assertNoError(t, err)

		want := State{Balances: map[Account]uint{"nick": 100_000, "alina": 1_000_000}}

		assertBalances(t, want, state)
	})
}

func assertBalances(t testing.TB, want, got State) {
	t.Helper()

	if !reflect.DeepEqual(want.Balances, got.Balances) {
		t.Errorf("expected %+v, but got %+v", want, got)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error, but got one %q", err)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("expected error, but didn't get one")
	}
}
