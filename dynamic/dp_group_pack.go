package main

import "fmt"

type pack struct {
	price, except int
}

// 动态规划，分组背包问题
func main() {
	var money, total int
	_, _ = fmt.Scan(&money, &total)
	money /= 10 // 10的倍数节省时间

	arr := make([][]pack, total+1)
	var price, p, q, except int

	for i := 1; i <= total; i++ {
		_, _ = fmt.Scan(&price, &p, &q)
		price /= 10 // 10的倍数节省时间
		except = price * p

		if q == 0 {
			if arr[i] == nil {
				arr[i] = make([]pack, 1)
			}
			arr[i][0] = pack{price, except}
		} else {
			if arr[q] == nil {
				arr[q] = make([]pack, 2)
				arr[q][1] = pack{price, except}
			} else {
				arr[q] = append(arr[q], pack{price, except})
			}
		}
	}

	dp := make([]int, money+1)
	// 计算1～total个物品的最大期望值，第n个物品依赖于n-1个物品的最大期望值
	for i := 1; i <= total; i++ {
		if arr[i] == nil {
			continue
		}
		for j := money; j >= arr[i][0].price; j-- {
			tmp := dp[j]

			// dp[j-arr[i][0].price] 表示购买该件商品后剩余钱所购买的最大期望值
			if tmp < dp[j-arr[i][0].price]+arr[i][0].except { // 主件
				tmp = dp[j-arr[i][0].price] + arr[i][0].except
			}

			// 有附件1
			if len(arr[i]) > 1 {
				price := arr[i][0].price + arr[i][1].price    // 主件 + 附件1 的单价
				except := arr[i][0].except + arr[i][1].except // 主件 + 附件1 的期望值

				if j >= price && tmp < dp[j-price]+except {
					tmp = dp[j-price] + except
				}
				// 有附件2
				if len(arr[i]) > 2 {
					// 主件 + 附件1 + 附件2
					price += arr[i][2].price
					except += arr[i][2].except
					if j >= price && tmp < dp[j-price]+except {
						tmp = dp[j-price] + except
					}

					//主件 + 附件2
					price -= arr[i][1].price
					except -= arr[i][1].except
					if j >= price && tmp < dp[j-price]+except {
						tmp = dp[j-price] + except
					}
				}
			}

			dp[j] = tmp
		}
	}
	fmt.Println(dp[money] * 10)
}
