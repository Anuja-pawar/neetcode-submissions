class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        sortedlist = sorted(nums)
        res = []
        for i, a in enumerate(sortedlist):
            if a > 0:
                break

            if i > 0 and a == sortedlist[i -1]:
                continue

            l, r = i+1, len(nums) - 1 
            while l < r :
                currsum = a + sortedlist[l] + sortedlist[r]
                if currsum > 0:
                    r -= 1
                elif currsum < 0:
                    l += 1
                else:
                    res.append([a,sortedlist[l],sortedlist[r]])
                    l += 1
                    r -= 1
                    while sortedlist[l] == sortedlist[l - 1] and l < r:
                        l += 1
        return res
                    


