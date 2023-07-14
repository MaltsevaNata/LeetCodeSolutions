const (
	BUY_OR_HOLD  = 1
	SELL_OR_WAIT = 0
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxProfit(prices []int, fee int) int {

	saved := make([][2]int, len(prices))
	saved[0][SELL_OR_WAIT] = 0
	saved[0][BUY_OR_HOLD] = 0 - prices[0] - fee

	for i := 1; i < len(prices); i++ {
		saved[i][SELL_OR_WAIT] = max(saved[i-1][SELL_OR_WAIT], saved[i-1][BUY_OR_HOLD]+prices[i])
		saved[i][BUY_OR_HOLD] = max(saved[i-1][BUY_OR_HOLD], saved[i-1][SELL_OR_WAIT]-prices[i]-fee)
	}
	return max(saved[len(prices)-1][SELL_OR_WAIT], saved[len(prices)-1][BUY_OR_HOLD])
}