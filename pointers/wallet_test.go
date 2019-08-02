package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("Expected %s, but got %s.", expected, actual)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}

		wallet.Withdraw(Bitcoin(5))

		actual := wallet.Balance()
		expected := Bitcoin(5)

		if actual != expected {
			t.Errorf("Expected %s, but got %s.", expected, actual)
		}
	})

}
