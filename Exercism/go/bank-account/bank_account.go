package account

import (
    "sync"
)

var mutex = &sync.Mutex{}

type Account struct {
    deposit int64
    status bool
}

func Open(initialDeposit int64) *Account {
    if initialDeposit < 0 {
        return nil
    }
	return &Account{initialDeposit, true}
}

func (ac *Account) Close() (payout int64, ok bool) {
	mutex.Lock()
    if ac.status == false {
    	mutex.Unlock()
        return 0, false
    }
	payout = ac.deposit
    ac.deposit = 0
    ac.status = false
    ok = true
    mutex.Unlock()
    return
}

func (ac *Account) Balance() (balance int64, ok bool) {
	mutex.Lock()
    if ac.status == false {
    	mutex.Unlock()
        return 0, false
    }
	balance = ac.deposit
    ok = true
    mutex.Unlock()
    return
}

func (ac *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	mutex.Lock()
    if ac.status == false || (amount < 0 && ac.deposit < (-1 * amount)){
    	mutex.Unlock()
        return 0, false
    }
	newBalance = ac.deposit + amount
	ac.deposit = newBalance
    ok = true
    mutex.Unlock()
    return
}