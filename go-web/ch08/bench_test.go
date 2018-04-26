package main

import (
	"testing"
	"time"
)

// benchmarking the decode function
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}
