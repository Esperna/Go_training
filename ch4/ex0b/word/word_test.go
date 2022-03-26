// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package word

import (
	"testing"
)

/*
func BenchmarkPrintSliceMacthUnefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintSliceMacthUnefficient()
	}
}
*/

func BenchmarkPrintSliceMacthEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintSliceMacthEfficient()
	}
}
