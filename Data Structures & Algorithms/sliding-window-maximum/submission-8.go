func maxSlidingWindow(nums []int, k int) []int {
    curr := make([]int, 0)

	res := make([]int, 0)
	for i:=0 ; i<len(nums); i++{
		if len(curr) == k{
			res = append(res, max(curr))
			curr = curr[1:]
		}
		curr = append(curr, nums[i])

	}
	res = append(res, max(curr))
	return res
}

func max(curr []int) int{
	maxN := curr[0]
	for _, n := range curr{
		if n > maxN{
			maxN = n
		}
	}
	return maxN
}