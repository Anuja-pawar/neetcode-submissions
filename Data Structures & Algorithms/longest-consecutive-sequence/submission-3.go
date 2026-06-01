func longestConsecutive(nums []int) int {
	longest := 0
	numsMap := make(map[int]bool)
	for _, n := range nums{
		numsMap[n] = true
	}
	for _, i := range nums{
		if !numsMap[i-1] {
			curr := 1
			for numsMap[i+1]{
				i++
				curr++
			}
			longest = max(curr, longest)
		}
	}
	return longest
}
