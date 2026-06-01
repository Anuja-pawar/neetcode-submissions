class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        upper = max(piles)
        lower = 1
        res = upper

        while lower <= upper:
            mid = (upper + lower) // 2
            hours = 0
            for p in piles:
                hours += math.ceil(p / mid)

            if hours <= h:
                res = min(res,mid)
                upper = mid -1
            else:
                lower = mid + 1
        return res
            


