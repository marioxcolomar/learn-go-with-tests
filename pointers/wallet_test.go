package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(250))

		got := wallet.Balance()

		want := Bitcoin(250)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(250)}

		wallet.Withdraw(Bitcoin(50))

		got := wallet.Balance()

		want := Bitcoin(200)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
