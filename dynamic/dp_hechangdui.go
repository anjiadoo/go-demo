/*
题目描述
计算最少出列多少位同学，使得剩下的同学排成合唱队形

说明：
N位同学站成一排，音乐老师要请其中的(N-K)位同学出列，使得剩下的K位同学排成合唱队形。
合唱队形是指这样的一种队形：设K位同学从左到右依次编号为1，2…，K，他们的身高分别为T1，T2，…，TK,
则他们的身高满足存在i（1<=i<=K）使得T1<T2<......<Ti-1<Ti>Ti+1>......>TK。

你的任务是，已知所有N位同学的身高，计算最少需要几位同学出列，可以使得剩下的同学排成合唱队形。

注意不允许改变队列元素的先后顺序
*/
package main

import "fmt"

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		dp_1 := make([]int, n)
		dp_2 := make([]int, n)

		arr := make([]int, n)
		for i := range arr {
			_, _ = fmt.Scan(&arr[i])
			dp_1[i] = 1
			dp_2[i] = 1
		}

		// 第一遍dp, 递增子序列长度
		for i := 0; i < n; i++ {
			for j := i - 1; j >= 0; j-- {
				if arr[i] > arr[j] && dp_1[j]+1 > dp_1[i] {
					dp_1[i] = dp_1[j] + 1
				}
			}
		}
		//fmt.Println(dp_1)

		// 第二遍dp，递增子序列长度
		reverse(arr)
		for i := 0; i < n; i++ {
			for j := i - 1; j >= 0; j-- {
				if arr[i] > arr[j] && dp_2[j]+1 > dp_2[i] {
					dp_2[i] = dp_2[j] + 1
				}
			}
		}
		//fmt.Println(dp_2)
		reverse(dp_2)

		ret := 0
		for i := range dp_1 {
			if dp_1[i]+dp_2[i] > ret {
				ret = dp_1[i] + dp_2[i]
			}
		}
		fmt.Println(n - ret + 1)
	}
}
