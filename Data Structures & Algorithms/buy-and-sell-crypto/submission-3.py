class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        l = 0
        r = l+1
        maxpft = 0
        while r < len(prices):
            if prices[l] < prices[r]:
                pft = prices[r] - prices[l]
                maxpft = max(maxpft, pft)
            else:
                l = r
            r +=1
        return maxpft
