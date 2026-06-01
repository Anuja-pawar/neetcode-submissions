func maxProfit(prices []int) int {
    mxProfit := 0 
    l, r := 0, 1

    for r < len(prices) {
        profit := prices[r] - prices[l]

        if prices[l] < prices[r]{
            if profit > mxProfit{
                mxProfit = profit
            }
        }else{
            l = r
        }
        r++
    }
    return mxProfit
}
