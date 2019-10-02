package main

import (
	"fmt"
	"math/rand"
)

func merge(r []int, begin, mid, end int) []int {
	var i = begin
	var j = mid + 1
	var k = begin
	var tmp = make([]int, len(r))

	for i <= mid && j <= end {
		if r[i] <= r[j] { /* 取r[i]与r[j]中较小者放入临时数组 */
			tmp[k] = r[i]
			k++
			i++
		} else {
			tmp[k] = r[j]
			k++
			j++
		}
	}
	for i <= mid { /* 对第一个子序列进行收尾处理 */
		tmp[k] = r[i]
		k++
		i++
	}
	for j <= end { /* 对第二个子序列进行收尾处理 */
		tmp[k] = r[j]
		k++
		j++
	}
	return tmp
}

func mergeSort(arr []int, begin, end int) {
	if begin == end {
		return
	} else {
		mid := (begin + end) / 2           /* 划分 */
		mergeSort(arr, begin, mid)         /* 求解子问题1 */
		mergeSort(arr, mid+1, end)         /* 求解子问题2 */
		tmp := merge(arr, begin, mid, end) /* 合并子序列 */
		for i := begin; i <= end; i++ { /* 将值传回原数组中 */
			arr[i] = tmp[i]
		}
	}
}

func main() {
	var size = 20
	arr := rand.Perm(size)
	fmt.Println(arr)

	mergeSort(arr, 0, size-1)
	fmt.Println(arr)
}
