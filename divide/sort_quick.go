package main

import (
	"fmt"
	"math/rand"
)

func partition(arr []int, fir, end int) int {
	i, j := fir, end
	for i < j {
		for i < j && arr[i] <= arr[j] { /* 右侧扫描 */
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		for i < j && arr[i] <= arr[j] { /* 左侧扫描 */
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}
	return i
}

/* 快速排序 */
func quickSort(arr []int, fir, end int) []int {
	var pivot int
	if fir < end {
		pivot = partition(arr, fir, end) /* 划分，确定轴值 */
		quickSort(arr, fir, pivot-1)     /* 求解子问题1 */
		quickSort(arr, pivot+1, end)     /* 求解子问题2 */
	}
	return arr
}

func main() {
	var size = 20
	arr := rand.Perm(size)
	fmt.Println(arr)

	quickSort(arr, 0, size-1)
	fmt.Println(arr)
}
