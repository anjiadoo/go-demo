package jz_offer

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
// 请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

func cuttingRope(n int) int {
	dp := make([]int, 3)

	dp[0] = 0
	dp[1] = 0
	dp[2] = 1

	for i := 3; i <= n; i++ {
		n1 := i / 2
		n2 := i - n1
		dp[(i % 3)] = getMaxValue([]int{n1 * n2, 3 * dp[(i-3)%3], 2 * dp[(i-2)%3]})
	}

	return getMaxValue(dp)
}

func getMaxValue(a []int) int {
	ret := a[0]
	for i := range a {
		if a[i] > ret {
			ret = a[i]
		}
	}
	return ret
}

func cuttingRope1(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	pow3 := func(n int) int {
		res := 1
		for i := 0; i < n; i++ {
			res = (res * 3) % 1000000007
		}
		return res
	}

	if n%3 == 0 {
		return pow3(n / 3)
	} else if n%3 == 1 {
		return pow3((n-3)/3) * 4 % 1000000007
	} else if n%3 == 2 {
		return pow3((n-2)/3) * 2 % 1000000007
	}
	return 0
}
