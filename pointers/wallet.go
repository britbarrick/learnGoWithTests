package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin is currency in wallet
type Bitcoin int

// Wallet will hold bitcoin
type Wallet struct {
	balance Bitcoin
}

// Stringer interface used to add BTC to each amount
type Stringer interface {
	String() string
}

// ErrInsufficientFunds used to ensure error text will not break program/tests if changed
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Deposit will add money to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance will provide running balance in wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw will take money from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
