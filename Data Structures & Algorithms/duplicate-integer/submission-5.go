func hasDuplicate(nums []int) bool {
    var seen = make(map[int]int, 0)
    for i, v := range nums{
        if _, ok := seen[v]; ok{
            return true
        }else{
            seen[v] = i
        }
    }
    return false
}
