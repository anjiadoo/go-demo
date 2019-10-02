package main

import (
	"fmt"
	"math/rand"
)

func insertSort(arr []int) []int {
	for i := 2; i < len(arr); i++ {
		arr[0] = arr[i]
		var j int
		for j = i - 1; arr[0] < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = arr[0]
	}
	arr[0] = -1 /* 观察哨 */
	return arr
}

func main() {
	var size int = 20
	arr := rand.Perm(size)

	arr[0] = -1 /* 观察哨 */
	fmt.Println(arr)

	arr = insertSort(arr)
	fmt.Println(arr)
}
