package main

import "fmt"

func dataTower(data [][]int) ([][]int, [][]int) {
	var n = len(data)
	var maxAdd = make([][]int, n)
	var path = make([][]int, n)

	/* path[i][j]: 表示在第i层第j个数塔的决策时选择的路径 */
	for i := 0; i < n; i++ {
		maxAdd[i] = make([]int, n) /* maxAdd: 存储每一步决策的结果--数值和 */
		path[i] = make([]int, n)   /* path: 存储每次决策所选数字在data中的下标 */
	}

	for j := 0; j < n; j++ {
		maxAdd[n-1][j] = data[n-1][j]
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if maxAdd[i+1][j] > maxAdd[i+1][j+1] {
				maxAdd[i][j] = data[i][j] + maxAdd[i+1][j]
				path[i][j] = j // 本次决策选择下标j的元素
			} else {
				maxAdd[i][j] = data[i][j] + maxAdd[i+1][j+1]
				path[i][j] = j + 1 // 本次决策选择下标j+1的元素
			}
		}
	}
	return path, maxAdd
}

/**
 * 有一个数塔，从数塔的顶层出发，在每一个节点可以选择向左走或者向右走，一直走到最底层
 * 要求找出一条路径，使得路劲上的数值和最大
 */

func main() {
	var data = [][]int{
		{80},
		{12, 15},
		{34, 38, 61},
		{81, 10, 51, 12},
		{16, 28, 18, 32, 41},
		{16, 41, 19, 37, 29, 53},
		{51, 34, 18, 10, 27, 73, 83},
	}
	path, max := dataTower(data)

	var j int
	for i := 0; i < len(data); i++ {
		if i == 0 {
			j = path[0][0] // 顶层决策是选择下一层列下标为path[0][0]的元素
			fmt.Printf("路径为：%d", data[0][0])
		} else {
			fmt.Printf("-->%d", data[i][j])
			j = path[i][j] // 本层决策是选择下一层列下标为path[i][j]的元素
		}
	}
	fmt.Println("\n最大数值和为：", max[0][0])

	for i := 0; i < len(max); i++ {
		fmt.Printf("\n第%d层最大值为: %v", i+1, max[i])
	}
}
