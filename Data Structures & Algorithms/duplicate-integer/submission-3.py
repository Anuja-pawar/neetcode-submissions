class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        copy = []
        for i in nums:
            if i not in copy:
                copy.append(i)
            else:
                return True
        return False

        