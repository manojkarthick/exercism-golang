package account

import "sync"

type Account struct {
	balance int64
	open    bool
	mutex   sync.RWMutex
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.open {
		return 0, false
	}
	payout = a.balance
	a.balance = 0
	a.open = false
	return payout, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	if !a.open {
		return 0, false
	}
	return a.balance, a.open
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.open || (a.balance+amount) < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, a.open
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
		open:    true,
	}
}
