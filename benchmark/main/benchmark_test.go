package main

import "testing"

var result int

func benchmarkFib1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib1(n)
	}

	result = r
}

func benchmarkFib2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib2(n)
	}

	result = r
}

func benchmarkFib3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib3(n)
	}

	result = r
}

func Benchmark30fibo1(b *testing.B) {
	benchmarkFib1(b, 30)
}

func Benchmark30fibo2(b *testing.B) {
	benchmarkFib2(b, 30)
}

func Benchmark30fibo3(b *testing.B) {
	benchmarkFib3(b, 30)
}

func Benchmark50fibo1(b *testing.B) {
	benchmarkFib1(b, 50)
}
func Benchmark50fibo2(b *testing.B) {
	benchmarkFib2(b, 50)
}
func Benchmark50fibo3(b *testing.B) {
	benchmarkFib3(b, 50)
}
