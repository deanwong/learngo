package main

import "fmt"

type Bank struct {
	accounts []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account2 > len(this.accounts) {
		return false
	}
	if this.Withdraw(account1, money) {
		return this.Deposit(account2, money)
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > len(this.accounts) {
		return false
	}
	this.accounts[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > len(this.accounts) {
		return false
	}
	if this.accounts[account-1] < money {
		return false
	}
	this.accounts[account-1] -= money
	return true
}

func main() {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	fmt.Println(bank.Withdraw(3, 10))    // true
	fmt.Println(bank.Transfer(5, 1, 20)) // true
	fmt.Println(bank.Deposit(5, 20))     // true
	fmt.Println(bank.Transfer(3, 4, 15)) // false
	fmt.Println(bank.Withdraw(10, 50))   // false
}
