package folkatech

import (
	"fmt"
	"sort"
)

func sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func addProfit(arr []int, val int) []int {
	i := len(arr)
	parent := (i - 1) / 2
	arr = append(arr, val)

	for arr[i] < arr[parent] {
		arr[i], arr[parent] = arr[parent], arr[i]
		i = parent
		parent = (i - 1) / 2
	}
	return arr
}

func MaxProfit(prices []int, chance int) int {
	if len(prices) < 2 {
		return 0
	}
	profits := []int{}
	buys := []int{}
	sells := []int{}
	max := 0
	sell := 0
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] > sell && prices[i] > buy {
			if sell == 0 {
				buys = append(buys, buy)
			}

			sell = prices[i]
		}
		if prices[i] < sell {
			sells = append(sells, sell)
			sell = 0
			buy = prices[i]
		}
		if prices[i] < buy {
			buy = prices[i]
		}
	}
	if sell != 0 {
		sells = append(sells, sell)
	}

	if chance < len(buys) /* or len(sells)*/ {
		for i := len(buys) - 1; /*or sells*/ i >= 0; i-- {
			a := sells[i] - buys[i]
			for j := i - 1; j >= chance; j-- {
				b := sells[i] - buys[j]
				if buys[j] < buys[i] && sells[i] > sells[j] {
					a = b
					i = j
				}
			}
			profits = append(profits, a)
		}
		sort.Slice(profits, func(i, j int) bool {
			return profits[i] > profits[j]
		})
		max = sum(profits[:chance])
	} else {
		for i := range buys /*or sells*/ {
			max += sells[i] - buys[i]
		}
	}
	return max
}

func MaxProfit4(prices []int, chance int) int {
	profit := 0
	buy, sell := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > buy && prices[i] > sell {
			sell = prices[i]
		}
		if prices[i] < sell && chance > 1 {
			profit += sell - buy
			buy = prices[i]
			sell = 0
			chance--
		}
		if i == len(prices)-1 {
			if sell != 0 {
				profit += sell - buy
			}
		}
	}
	// if len(profits) > chance {
	// 	sort.Slice(profits, func(i, j int) bool {
	// 		return profits[i] > profits[j]
	// 	})
	// 	maxProfit = sum(profits[:chance])
	// } else {
	// 	maxProfit = sum(profits)
	// }

	return profit
}

func MaxProfit3(prices []int, chance int) int {
	profits := []int{}
	buy, sell := prices[0], 0
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > buy && prices[i] > sell {
			sell = prices[i]
		}
		if prices[i] < sell {
			profits = append(profits, sell-buy)
			buy = prices[i]
			sell = 0
		}
		if i == len(prices)-1 {
			if sell != 0 {
				profits = append(profits, sell-buy)
			}
		}
	}
	if len(profits) > chance {
		sort.Slice(profits, func(i, j int) bool {
			return profits[i] > profits[j]
		})
		maxProfit = sum(profits[:chance])
	} else {
		maxProfit = sum(profits)
	}

	return maxProfit
}
func MaxProfit2(prices []int, chance int) int {

	buy, sell, profit := prices[0], 0, 0

	for i := 1; i < len(prices); i++ {
		fmt.Println(i)
		if buy < prices[i] {
			sell = prices[i]
		}
		if i < len(prices)-1 {
			if prices[i+1] <= sell && chance > 0 {
				profit += sell - buy
				chance--
				if i+1 < len(prices)-1 {
					buy = prices[i+1]
				} else {
					chance = 0
				}
			}
		}
	}
	if chance > 0 {
		profit += sell - buy
	}
	fmt.Println(buy, sell)
	return profit
}
