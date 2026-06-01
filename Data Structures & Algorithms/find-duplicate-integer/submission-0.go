func findDuplicate(nums []int) int {
    countN := make(map[int]int)
    for _, i := range nums{
        countN[i]++
        if countN[i] > 1{
            return i
        }
    }
    return 0
}
