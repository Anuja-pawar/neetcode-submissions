func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 0{
        return []int{}
    }
    curr := make([]int, 0)
	res := make([]int, 0)
	
	for i:=0 ; i<len(nums); i++{
		//maintain curr window
		if len(curr) >0 && curr[0] <= i - k{
			curr = curr[1:]
		}
		//keep max element in curr
		for len(curr) >0 && nums[curr[len(curr)-1]] < nums[i]{
			curr = curr[:len(curr)-1]
		}
		//append current
		curr = append(curr, i)
		//start adding res once we hit the first window
		if i >= k - 1{
			res = append(res, nums[curr[0]])
		}	
	}
	return res
}
