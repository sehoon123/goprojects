package main

import "testing"

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(40)
	}
}
func BenchmarkFibonacci3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci3(40)
	}
}
func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(40)
	}
}
