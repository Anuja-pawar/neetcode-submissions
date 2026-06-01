func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    var result []int
    for i := range nums{
        diff := target - nums[i]
        for key, value := range numsMap {
            if value == diff {
                result = append(result, key)
                result = append(result, i)
                return result
            }
	    }
        numsMap[i] = nums[i]
    }
    return result
}
