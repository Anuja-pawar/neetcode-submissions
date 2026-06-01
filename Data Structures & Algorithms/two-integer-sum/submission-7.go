func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, n := range nums{
		diff := target - n
		if j, ok := numsMap[diff]; ok {
			return []int{j, i}
		}
		numsMap[n] = i
	}
	return []int{}
}
