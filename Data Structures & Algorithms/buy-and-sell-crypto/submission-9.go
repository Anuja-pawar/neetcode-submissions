func maxProfit(prices []int) int {
	buy := 0
	maxProfit := 0
	for i := 1; i<len(prices); i++{
		currProfit := prices[i] - prices[buy]
		if currProfit > maxProfit{
			maxProfit = currProfit
		}else if prices[i] < prices[buy]{
			buy = i
		}
	}
	return maxProfit

}
