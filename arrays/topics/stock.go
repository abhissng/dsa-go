package topics

func StockBasic(arr []int) int {
	return maxProfit(arr, 0, len(arr)-1)
}
func maxProfit(price []int, start, end int) int {
	if end <= start {
		return 0
	}
	profit := 0

	for i := start; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			if price[j] > price[i] {
				currentProfit := price[j] - price[i] + maxProfit(price, start, i-1) + maxProfit(price, j+1, end)
				profit = max(profit, currentProfit)
			}
		}
	}
	return profit
}

func StockBetter(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
