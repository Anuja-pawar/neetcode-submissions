class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        seen = []
        for n in nums:
            print(n)
            if n not in seen:
                seen.append(n)
            else:
                return True
        return False

        