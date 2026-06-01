func maxProfit(prices []int) int {
    mxProfit := 0 
    currProfit := 0
    for i := range prices{
        l := i
        r := i+1
        for l < r && r < len(prices) {
            buy := prices[l]
            sell := prices[r]
            currProfit = 0
            fmt.Println(buy)
            fmt.Println(sell)
            fmt.Println(currProfit)
            if buy > sell{
                r++
            }else{
                currProfit = sell - buy
                r++
            }
            if currProfit > mxProfit{
                mxProfit = currProfit
            }
            fmt.Println("mx ",mxProfit)
        }
    }
    if mxProfit < 0{
        return 0
    }
    return mxProfit
}
