func maxProfit(prices []int) int {
    maxProfit := 0
    minStock := prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i] < minStock {
            minStock = prices[i]
        } else if (prices[i] - minStock) > maxProfit {
            maxProfit = prices[i] - minStock
        }
    }
    return maxProfit
}
