package jz_offer

// 实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	power := 1.0
	if n > 0 {
		power = myPow(x, n/2)
		return power * power * myPow(x, n%2)
	}
	if n < 0 {
		n = -n
		power = myPow(x, n/2)
		power = power * power * myPow(x, n%2)
		return 1 / power
	}
	return power
}
