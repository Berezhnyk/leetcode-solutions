package maxprofit

func maxProfit(prices []int) int {
	lp := len(prices)
	if lp == 0 || lp == 1 {
		return 0
	}

	maxProfit := 0
	buyLowest := prices[0]

	for _, price := range prices {
		if price < buyLowest {
			buyLowest = price
			continue
		}
		profit := price - buyLowest
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
