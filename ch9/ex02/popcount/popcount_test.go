// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"ch9/ex02/popcount"
	"sync"
	"testing"
)

func TestLoadpc1time(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popcount.PopCount(0xFFFFFFFFFFFFFFFF)
		}()
	}
	wg.Wait()
	if popcount.NumberOfcallLoadpc != 1 {
		t.Errorf("expected numberOfcallLoadpc: %d actual: %d", 1, popcount.NumberOfcallLoadpc)
	}
}
