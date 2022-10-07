package database

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadGenesis(t *testing.T) {
	cwd, _ := os.Getwd()

	t.Run("load statefull genesis file", func(t *testing.T) {
		genesis, err := loadGensis(filepath.Join(cwd, "test_files", "genesis_statefull.json"))

		if err != nil {
			t.Fatalf("expected no error, but got one %q", err)
		}

		want := Genesis{
			Balances: map[Account]uint{
				"nick": 1_000_000,
			},
		}

		if !reflect.DeepEqual(genesis, want) {
			t.Errorf("expected %v, but got %v", want, genesis)
		}
	})

	t.Run("load empty genesis file", func(t *testing.T) {
		genesis, err := loadGensis(filepath.Join(cwd, "test_files", "genesis_empty.json"))

		if err != nil {
			t.Fatalf("expected no error, but got one %q", err)
		}

		want := Genesis{
			Balances: map[Account]uint{},
		}

		if !reflect.DeepEqual(genesis, want) {
			t.Errorf("expected %v, but got %v", want, genesis)
		}
	})
}
