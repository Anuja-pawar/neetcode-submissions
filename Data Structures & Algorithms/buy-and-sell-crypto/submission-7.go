func maxProfit(prices []int) int {
    buy := 0
    sell := 1
    maxProfit := 0
    for buy < sell && sell < len(prices){
        if prices[sell] > prices[buy]{
            profit := prices[sell] - prices[buy]
            maxProfit = max(profit, maxProfit)
        }else{
            buy = sell
        }
        sell++
    }
    return maxProfit
}
