// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"ch9/ex01/bank"
	"fmt"
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
}
