package jz_offer

//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

func numWays(n int) int {
	if n < 1 {
		return 1
	}

	// f(n) = f(n-1)+f(n-2)

	way, f0, f1 := 0, 1, 1
	for i := 0; i < n-1; i++ {
		way = (f0 + f1) % 1000000007
		f0, f1 = f1, way
	}
	return way
}
