package main

import "fmt"

type pair struct {
	fir int
	sec int
}

func stoneGame(stones []int) int {

	/**
	 * 状态有三个：开始索引i，结束索引j，当前轮到的人。
	 * dpt[i][j](fir or sec)，其中 0 <= i < len(stones)， i <= j < len(stones)；
	 * dpt[i][j].fir 表示，对于stones[i...j]这部分石头堆，先手能获得的最高分数；
	 * dpt[i][j].sec 表示，对于stones[i...j]这部分石头堆，后手能获得的最高分数；
	 */
	var dpt = make([][]pair, len(stones))

	//初始化dp table
	for i := 0; i < len(stones); i++ {
		dpt[i] = make([]pair, len(stones))
		for j := 0; j < len(stones); j++ {
			dpt[i][j] = pair{0, 0}
			dpt[i][i].fir = stones[i] //只有一堆石头时，先手获得
		}
	}

	//斜着遍历数组
	for k := 2; k <= len(stones); k++ {
		for i := 0; i <= len(stones)-k; i++ {
			j := k + i - 1

			/* 先手(A)取走一堆石子后，轮到后手(B)，相对于现在的后手(B)，这时的先手(A)变成了后手 */
			/* 所以，left 的含义是：A选取的左边石子 + 剩下的石子中A作为后手的最大分数 */

			//先手选择最左or最右的分数
			left := stones[i] + dpt[i+1][j].sec
			right := stones[j] + dpt[i][j-1].sec

			//状态转移方程
			if left > right {
				dpt[i][j].fir = left
				dpt[i][j].sec = dpt[i+1][j].fir
			} else {
				dpt[i][j].fir = right
				dpt[i][j].sec = dpt[i][j-1].fir
			}
		}
	}

	for i := 0; i < len(stones); i++ {
		fmt.Println(dpt[i])
	}
	res := dpt[0][len(stones)-1]
	return res.fir - res.sec
}

func main() {
	var stones = []int{30, 59, 1, 24, 54, 88, 2, 12, 90}
	fmt.Println("先后手分数之差：", stoneGame(stones))
}
