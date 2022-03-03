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
