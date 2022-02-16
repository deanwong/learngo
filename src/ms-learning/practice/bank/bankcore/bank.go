package bank

import (
	"errors"
	"fmt"
)

type Accountable interface {
	Statement() string
}

// func Statement(a Accountable) string {
// 	return a.Statement()
// }

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount should be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount should be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("balance not enough")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%d - %s - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(amount float64, target *Account) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	if err := target.Deposit(amount); err != nil {
		return err
	}
	return nil
}
