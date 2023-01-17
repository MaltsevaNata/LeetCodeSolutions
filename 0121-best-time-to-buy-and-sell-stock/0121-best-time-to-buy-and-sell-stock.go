func maxProfit(prices []int) int {
	length := len(prices)
	profit := 0
	maxProfit := 0
	minValue := prices[0]
	maxValue := 0

	foundPair := true

	for i := 1; i < length; i++ {
		if prices[i]-prices[i-1] < 0 {
			if prices[i] < minValue {
				foundPair = false
				minValue = prices[i]
			}
		} else {
			if prices[i] > maxValue || !foundPair {
				maxValue = prices[i]
				profit = maxValue - minValue
				if profit > maxProfit {
					maxProfit = profit
				}
				foundPair = true
			}
		}
	}
	if maxProfit < 0 {
		return 0
	}
	return maxProfit
}