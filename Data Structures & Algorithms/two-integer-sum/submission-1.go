func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i, n := range nums{
        diff := target - n
        if ind, ok := numsMap[diff]; ok{
            return []int{ind, i}
        }
        numsMap[n] = i
    }
    return []int{}
}
