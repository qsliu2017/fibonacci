package fibonacci

func Fibonacci1(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fibonacci1(n-1) + Fibonacci1(n-2)
	}
}

func Fibonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fib := make([]int, n+1)
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func Fibonacci3(n int) int {
	if n == 0 {
		return 0
	}
	ans, _, _, _ := matrixN(n - 1)
	return ans
}

// return result of
// | 1 1 | ^ n
// | 1 0 |
func matrixN(n int) (z11, z12, z21, z22 int) {
	if n == 1 {
		return 1, 1, 1, 0
	}
	x11, x12, x21, x22 := matrixN(n / 2)
	y11, y12, y21, y22 := matrixDot(x11, x12, x21, x22, x11, x12, x21, x22)
	if n%2 == 0 {
		return y11, y12, y21, y22
	} else {
		return matrixDot(y11, y12, y21, y22, 1, 1, 1, 0)
	}
}

func matrixDot(x11, x12, x21, x22, y11, y12, y21, y22 int) (z11, z12, z21, z22 int) {
	return x11*y11 + x12*y21, x11*y12 + x12*y22, x21*y11 + x22*y21, x21*y12 + x22*y22
}
