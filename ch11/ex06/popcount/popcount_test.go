// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"ch2/ex05/popcount"
	"testing"
)

// -- Alternative implementations --

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func TestPopCount(t *testing.T) {
	var tests = []struct {
		want  int
		given uint64
		f     func(uint64) int
	}{
		{32, 0x1234567890ABCDEF, popcount.PopCount},
		{32, 0x1234567890ABCDEF, popcount.PopCountByShift},
		{32, 0x1234567890ABCDEF, popcount.PopCountByLoop},
		{32, 0x1234567890ABCDEF, BitCount},
		{32, 0x1234567890ABCDEF, popcount.PopCountByClear},
		{32, 0x1234567890ABCDEF, PopCountByClearing},
		{32, 0x1234567890ABCDEF, PopCountByShifting},
	}
	for _, test := range tests {
		actual := test.f(test.given)
		if actual != test.want {
			t.Errorf("(%d): want %d, actual %d", test.given, test.want, actual)
		}
	}

}

// -- Benchmarks --

func BenchmarkPopCount16digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShift16digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShift(0x1234567890ABCDEF)
	}

}

func BenchmarkPopCountByLoop16digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount16digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClear16digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClear(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing16digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting16digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount8digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x12345678)
	}
}

func BenchmarkPopCountByShift8digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShift(0x12345678)
	}

}

func BenchmarkPopCountByLoop8digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoop(0x12345678)
	}
}

func BenchmarkBitCount8digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x12345678)
	}
}

func BenchmarkPopCountByClear8digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClear(0x12345678)
	}
}

func BenchmarkPopCountByClearing8digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x12345678)
	}
}

func BenchmarkPopCountByShifting8digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x12345678)
	}
}

func BenchmarkPopCount4digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234)
	}
}

func BenchmarkPopCountByShift4digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShift(0x1234)
	}

}

func BenchmarkPopCountByLoop4digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoop(0x1234)
	}
}

func BenchmarkBitCount4digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234)
	}
}

func BenchmarkPopCountByClear4digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClear(0x1234)
	}
}

func BenchmarkPopCountByClearing4digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234)
	}
}

func BenchmarkPopCountByShifting4digits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234)
	}
}

// Go 1.6, 2.67GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         6.30 ns/op
// BenchmarkBitCount-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// Go 1.6, 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         7.52 ns/op
// BenchmarkBitCount-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op
//
// Go 1.7, 3.5GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-12                 2000000000        0.28 ns/op
// BenchmarkBitCount-12                 2000000000        0.27 ns/op
// BenchmarkPopCountByClearing-12       100000000        18.5 ns/op
// BenchmarkPopCountByShifting-12       20000000         70.1 ns/op