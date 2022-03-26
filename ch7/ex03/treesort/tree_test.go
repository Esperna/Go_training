package treesort

import (
	"testing"
)

func TestString(t *testing.T) {
	var tests = []struct {
		values   []int
		expected string
	}{
		{[]int{1}, "{1}"},
		{[]int{1, 2}, "{1 2}"},
		{[]int{3, 2, 1}, "{1 2 3}"},
		{[]int{3, 5, 9, 0, 4, 2, 7, 1, 6, 8}, "{0 1 2 3 4 5 6 7 8 9}"},
	}
	for _, test := range tests {
		var root *tree
		for _, v := range test.values {
			root = add(root, v)
		}
		if root.String() != test.expected {
			t.Errorf("String() is not expected. expected is %s String() is %s", test.expected, root.String())
		}
	}
}
