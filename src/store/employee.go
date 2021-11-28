package store

import (
	"errors"
	"fmt"
)

type Account struct {
	first string
	last  string
}

func (account *Account) ChangeName(newName string) {
	account.first = newName
}

type Employee struct {
	credits float64
	Account
}

func NewEmployee(first, last string, credits float64) (*Employee, error) {
	return &Employee{credits, Account{first, last}}, nil
}

func (employee *Employee) AddCredits(credits float64) float64 {
	employee.credits += credits
	return employee.credits
}

func (employee *Employee) RemoveCredits(credits float64) (float64, error) {
	if credits > 0.0 {
		if credits <= employee.credits {
			employee.credits -= credits
			return employee.credits, nil
		}
		return 0.0, errors.New("Larger then he/she have")
	}
	return 0.0, errors.New("Negative number")
}

func (employee *Employee) CheckCredits() float64 {
	return employee.credits
}

func (account *Account) String() string {
	return fmt.Sprintf("%v %v", account.first, account.last)
}
