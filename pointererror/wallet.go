package pointererror

import (
	"errors"
	"fmt"
)

// Bitcoin 传说中的比特币
type Bitcoin int

// Wallet : 钱包
type Wallet struct {
	balance Bitcoin
}

// Deposit : 存款
func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

// Balance : 获取余额
func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

// ErrInsufficientFundsError :
var ErrInsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// Withdraw : 取款
func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficientFundsError
	}

	wallet.balance -= amount
	return nil
}

// 实现 fmt 包的 Stringer 接口的 String 方法，定义 %s 的打印方式
// wallet_test.go:18: got 10 BTC want 20 BTC
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
