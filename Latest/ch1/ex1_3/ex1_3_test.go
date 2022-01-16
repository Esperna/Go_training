package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) { Echo1() }
func BenchmarkEcho2(b *testing.B) { Echo2() }
