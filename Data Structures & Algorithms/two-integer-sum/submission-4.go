func twoSum(nums []int, target int) []int {
    for i, a := range nums{
		b := target - a
		for j := i+1; j < len(nums); j++{
			if nums[j] == b{
				return []int{i, j}
			}
		}
	}
	return []int{}
}
