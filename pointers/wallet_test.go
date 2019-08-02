package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{10}
		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds.Error())
	})

}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	actual := wallet.Balance()

	if actual != expected {
		t.Errorf("Expected %s, but got %s.", expected, actual)
	}
}

func assertError(t *testing.T, actual error, expected string) {
	t.Helper()
	if actual == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if actual.Error() != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
