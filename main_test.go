package main

import "testing"

func BenchmarkQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Query()
	}
}
