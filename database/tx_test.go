package database

import (
	"testing"
)

func TestIsReward(t *testing.T) {
	t.Run("tx is a reward", func(t *testing.T) {
		tx := Tx{From: "nick", To: "alina", Value: 1000, Data: "reward"}

		got := tx.isReward()
		want := true

		if got != want {
			t.Errorf("got %t, but want %t, given %#v", got, want, tx)
		}
	})

	t.Run("tx isn't a reward", func(t *testing.T) {
		tx := Tx{From: "nick", To: "alina", Value: 1000, Data: "other"}

		got := tx.isReward()
		want := false

		if got != want {
			t.Errorf("got %t, but want %t, given %#v", got, want, tx)
		}
	})
}
