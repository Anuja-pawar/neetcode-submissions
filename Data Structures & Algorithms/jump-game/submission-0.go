func canJump(nums []int) bool {
	farthest := 0
	for i:=0; i<len(nums); i++{
		if farthest < i{
			return false
		}
		if i + nums[i] > farthest{
			farthest = i + nums[i]
		}
		if farthest >= len(nums) - 1{
			return true
		}
	}
	return false
}
