// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"ch9/ex01/bank"
	"fmt"
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	type withdrawResult struct {
		isSuccess bool
		balance   int
	}
	var tests = []struct {
		expected withdrawResult
		given    int
	}{
		{withdrawResult{false, 300}, 400},
		{withdrawResult{true, 200}, 100},
	}
	for _, test := range tests {
		var isSuccess bool
		// Bob
		go func() {
			isSuccess = bank.Withdraw(test.given)
			done <- struct{}{}
		}()

		// Wait for both transactions.
		<-done
		if isSuccess != test.expected.isSuccess {
			t.Errorf("Withdraw returns %t want %t", isSuccess, test.expected.isSuccess)
		}
		if got, want := bank.Balance(), test.expected.balance; got != want {
			t.Errorf("Balance = %d, want %d", got, want)
		}
	}
	//tear down
	bank.Withdraw(bank.Balance())
	if bank.Balance() != 0 {
		t.Errorf("Balance = %d, want 0", bank.Balance())
	}
}

func TestBankConcurrentWithdraw(t *testing.T) {
	result := make(chan bool)
	const numOfTest = 10000
	for i := 0; i < numOfTest; i++ {
		ready := make(chan struct{})
		bank.Withdraw(bank.Balance())
		if bank.Balance() != 0 {
			t.Errorf("Balance = %d, want 0", bank.Balance())
		}
		bank.Deposit(100)
		const numOfGoRoutine = 10
		var n sync.WaitGroup
		for i := 0; i < numOfGoRoutine; i++ {
			n.Add(1)
			go func() {
				n.Done()
				<-ready
				result <- bank.Withdraw(60)
			}()
		}
		n.Wait()
		close(ready)
		successCount := 0
		for i := 0; i < numOfGoRoutine; i++ {
			if <-result {
				successCount++
			}
		}
		const want = 40
		if bank.Balance() != want {
			t.Errorf("balance is %d, want %d", bank.Balance(), want)
		}
		if successCount != 1 {
			t.Errorf("deposit success count is %d, want 1", successCount)
		}
	}
}

func TestBankConcurrentDeposit(t *testing.T) {
	done := make(chan struct{})
	const numOfTest = 10000
	for i := 0; i < numOfTest; i++ {
		ready := make(chan struct{})
		bank.Withdraw(bank.Balance())
		if bank.Balance() != 0 {
			t.Errorf("Balance = %d, want 0", bank.Balance())
		}
		const numOfGoRoutine = 10
		var n sync.WaitGroup
		for i := 0; i < numOfGoRoutine; i++ {
			n.Add(1)
			go func() {
				n.Done()
				<-ready
				bank.Deposit(100)
				done <- struct{}{}
			}()
		}
		n.Wait()
		close(ready)
		for i := 0; i < numOfGoRoutine; i++ {
			<-done
		}
		const want = 100 * numOfGoRoutine
		if bank.Balance() != want {
			t.Errorf("balance is %d, want %d", bank.Balance(), want)
		}
	}
}
