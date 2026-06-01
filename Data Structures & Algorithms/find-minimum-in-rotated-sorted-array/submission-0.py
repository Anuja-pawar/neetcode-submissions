class Solution:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums)-1
        minimum = nums[l]
        while l <= r:
            mid = (l+r)//2
            if nums[l] < nums[r]:
                minimum = min(nums[l], minimum)
                break
            minimum = min(nums[mid], minimum)
            if nums[l] <= nums[mid]:
                l = mid + 1
            elif nums[mid] < nums[r]:
                r = mid - 1
        return minimum

        