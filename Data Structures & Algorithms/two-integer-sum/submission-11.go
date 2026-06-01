func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i, v := range nums{
        numsMap[v] = i
    }

    for i, v := range nums{
        rem := target - v
        if j, ok := numsMap[rem]; ok && i != j{
            return []int{i, j}
        }
    }
    return []int{}
}
