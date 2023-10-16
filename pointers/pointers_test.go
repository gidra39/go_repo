package pointers

import "testing"

type Wallet struct {
	balance int
}

type Bitcoin int

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += int(amount)
}

func (w *Wallet) Balance() interface{} {
	return w.balance
}

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
