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
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(5))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	actual := wallet.Balance()

	if actual != expected {
		t.Errorf("Expected %s, but got %s.", expected, actual)
	}
}

func assertNoError(t *testing.T, actual error) {
	t.Helper()
	if actual != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t *testing.T, actual error, expected error) {
	t.Helper()
	if actual == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
