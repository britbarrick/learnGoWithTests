package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	actual := wallet.Balance()
	expected := Bitcoin(10)

	if actual != expected {
		t.Errorf("Expected %d, but got %d.", expected, actual)
	}
}
