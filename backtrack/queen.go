package main

import (
	"fmt"
	"math"
)

func queen(n int) {
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = -1
	}

	var k = int(0)
	for k >= 0 {
		arr[k]++
		for arr[k] < n && place(arr, k) == 1 {
			arr[k]++
		}

		if arr[k] < n && k == n-1 {
			for i := 0; i < len(arr); i++ {
				fmt.Println(arr[i] + 1)
			}
			return
		}

		if arr[k] < n && k < n-1 {
			k = k + 1
		} else {
			k--
			arr[k] = -1
		}
	}
}

func place(arr []int, k int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[k] || math.Abs(float64(i-k)) == math.Abs(float64(arr[i]-arr[k])) {
			return 1
		}
	}
	return 0
}

func main() {
	var size = int(4)
	queen(size)
}
