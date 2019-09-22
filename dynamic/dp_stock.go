package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func singleMaxProfit(prices []int) (int, int) {
	if len(prices) <= 1 {
		return -1, 0
	}
	idx, buy, sell := -1, -prices[0], 0
	for i := 0; i < len(prices); i++ {
		buy = max(buy, -prices[i])
		if prices[i]+buy > sell {
			idx = i
		}
		sell = max(sell, prices[i]+buy)
	}
	return idx + 1, sell
}

func secondMaxProfit(prices []int) (int) {
	if len(prices) <= 1 {
		return 0
	}
	fstBuy, fstSell, secBuy, secSell := -prices[0], 0, 0, 0
	for i := 0; i < len(prices); i++ {
		fstBuy = max(fstBuy, -prices[i])
		fstSell = max(fstSell, prices[i]+fstBuy)
		secBuy = max(secBuy, fstSell-prices[i])
		secSell = max(secSell, prices[i]+secBuy)
	}
	return secSell
}

func main() {
	var prices = []int{7, 1, 5, 3, 60, 4}
	fmt.Println(singleMaxProfit(prices))

	var prices1 = []int{3, 3, 9, 0, 3, 1, 4, 5, 0, 6}
	fmt.Println(secondMaxProfit(prices1))
}
