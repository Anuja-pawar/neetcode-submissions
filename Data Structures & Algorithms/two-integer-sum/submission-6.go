func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, n := range nums{
		numsMap[n] = i
	}

	for i, n1 := range nums{
		n2 := target - n1
		if val, ok := numsMap[n2]; ok && i != val{
			return []int{i, val}
		}
	}
	return []int{}
}
