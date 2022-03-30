// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawInfo struct {
	amount int
	result chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdrawInfo)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			if balance >= withdraw.amount {
				balance -= withdraw.amount
				withdraw.result <- true
			} else {
				withdraw.result <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func Withdraw(amount int) bool {
	result := make(chan bool)
	withdraws <- withdrawInfo{amount, result}
	return <-result
}

//!-
