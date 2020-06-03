package jz_offer

//写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：
//
//F(0) = 0,   F(1) = 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

func fib(n int) int {
	if n < 2 {
		return n
	}

	sum, pre, curr := 0, 0, 1
	for i := 0; i < n-1; i++ {
		sum = (pre + curr) % 1000000007
		pre, curr = curr, sum
	}
	return sum
}

func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}
