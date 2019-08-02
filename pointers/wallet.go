package pointers

// Bitcoin is currency in wallet
type Bitcoin int

// Wallet will hold bitcoin
type Wallet struct {
	balance Bitcoin
}

// Deposit will add money to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance will provide running balance in wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
