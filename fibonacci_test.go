package fibonacci_test

import (
	"fibonacci"
	"testing"
)

func TestFibonacci(t *testing.T) {
	N := 100
	for i := 2; i < N; i++ {
		// Do run this. Too slow.
		// ans1 := fibonacci.Fibonacci1(i)
		ans2 := fibonacci.Fibonacci2(i)
		ans3 := fibonacci.Fibonacci3(i)
		if ans2 != ans3 {
			t.Errorf("error when i=%d, ans2=%d, ans3=%d", i, ans2, ans3)
		}
	}
}

// Do run this. Too slow.
func BenchmarkFibonacci1(b *testing.B) {
	fibonacci.Fibonacci1(b.N)
}

func BenchmarkFibonacci2(b *testing.B) {
	fibonacci.Fibonacci2(b.N)
}

func BenchmarkFibonacci3(b *testing.B) {
	fibonacci.Fibonacci3(b.N)
}
