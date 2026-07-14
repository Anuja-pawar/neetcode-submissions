func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i:=0; i < len(nums); i++{
        a := nums[i]
        if a > 0 {
            break
        }
        if i > 0 && a == nums[i-1]{
            continue
        }
		l := i+1
		r := len(nums)-1
		for l < r{
			currSum := a + nums[l] + nums[r]
			if currSum > 0{
				r--
			}else if currSum < 0{
				l++
			}else{
				res = append(res, []int{a, nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
			}
		}
	}
	return res
}
