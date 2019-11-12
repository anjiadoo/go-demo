package main

import (
	"fmt"
	"math"
)

func primeCircle(n int) {
	var arr = []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, 0)
	}
	arr[0] = 1
	var k = 1
	for k >= 1 {
		fmt.Println(k, arr, )
		arr[k] = arr[k] + 1

		for arr[k] <= n {
			if check(arr, k) == 1 {
				break
			} else {
				arr[k] = arr[k] + 1
			}
		}

		/* 求解完毕，输出解 */
		if arr[k] <= n && k == n-1 {
			for i := 0; i < n; i++ {
				fmt.Printf("%d, ", arr[i])
			}
			return
		}

		if arr[k] <= n && k < n-1 {
			k++ // 填写下一个位置
		} else {
			k-- // 回溯
			arr[k] = 0
		}
	}
}

/* 判断位置k的填写是否满足约束条件 */
func check(arr []int, k int) int {
	var flag = int(0)
	for i := 0; i < k; i++ { // 判断是否重复
		if arr[i] == arr[k] {
			return 0
		}
	}
	flag = prime(arr[k] + arr[k-1]) // 判断相邻之和是否为素数
	if flag == 1 && k == len(arr)-1 {
		flag = prime(arr[k] + arr[0]) // 判断首尾之和是否为素数
	}
	return flag
}

/* 判断整数x是否为素数 */
func prime(x int) int {
	k := math.Sqrt(float64(x))
	n := int(k)

	for i := 2; i <= n; i++ {
		if x%i == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	var size = int(20)
	primeCircle(size)
}
