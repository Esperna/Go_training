// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 359.

// Package equal provides a deep equivalence relation for arbitrary values.
package cyclic

import (
	"reflect"
	"unsafe"
)

func IsCyclic(x interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return isCyclic(reflect.ValueOf(x), seen)
}

func isCyclic(x reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !x.IsValid() {
		return false
	}
	kind := x.Kind()
	if x.CanAddr() && kind != reflect.Struct && kind != reflect.Array {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		if seen[xptr] {
			return true // already seen
		}
		seen[xptr] = true
	}
	switch kind {
	case reflect.Ptr, reflect.Interface:
		return isCyclic(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCyclic(x.Index(i), seen) {
				return true
			}
		}
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCyclic(x.Field(i), seen) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCyclic(x.MapIndex(k), seen) {
				return true
			}
		}
	}
	return false
}
