package main

import (
	"fmt"
	"math/rand"
)

func siftHeap(arr []int, k, n int) {
	/* i为要筛选的节点，j为i的左孩子 */
	var i int = k
	var j int = 2*i + 1

	/* 筛选还没有进行到叶子节点 */
	for j < n {

		/* 比较i的左右孩子，j为较大者 */
		if j < n-1 && arr[j] < arr[j+1] {
			j++
		}
		if arr[i] > arr[j] {
			break
		} else { // 将被筛选节点与j交换
			arr[i], arr[j] = arr[j], arr[i]
			i = j; // 被筛选节点位于原来节点j的位置
			j = 2*i + 1
		}
	}
}

func heapSort(arr []int, n int) []int {
	var ret = make([]int, n)

	/* 初始建堆，最后一个分支的下标是(n-1)/2 */
	for i := (n - 1) / 2; i >= 0; i-- {
		siftHeap(arr, i, n)
	}

	/* 重复执行移走堆顶及重建堆的操作 */
	for i := 1; i <= n-1; i++ {
		ret[i-1] = arr[0]

		arr[0], arr[n-i] = arr[n-i], arr[0]
		siftHeap(arr, 0, n-i) // 只需要调整根节点
	}
	/* ret: 降序; arr: 升序 */
	return arr
}

/**
 * [12 4 2 13 10 0 19 11 7 5 15 18 9 14 6 8 1 16 17 3]
 * [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
 */
func main() {
	var size int = 20
	arr := rand.Perm(size)
	fmt.Println(arr)

	arr = heapSort(arr, size)
	fmt.Println(arr)
}
