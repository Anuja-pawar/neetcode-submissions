class Solution:
    def search(self, nums: List[int], target: int) -> int:
        start = 0
        end = len(nums) - 1
        mid = start + (end - start)//2

        while start <= end and end >= 0:
            if target > nums[mid]:
                start = mid + 1
            elif target < nums[mid]:
                end = mid - 1
            else:
                return mid
            mid = start + (end - start)//2
        
        return -1