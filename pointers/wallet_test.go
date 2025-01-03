package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(250))
		assertBalance(t, wallet, Bitcoin(250))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(250)}
		err := wallet.Withdraw(Bitcoin(50))

		assertError(t, err, nil)
		assertBalance(t, wallet, Bitcoin(200))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(150))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(50))

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
