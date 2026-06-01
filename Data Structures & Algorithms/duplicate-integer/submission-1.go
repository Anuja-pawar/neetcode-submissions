func hasDuplicate(nums []int) bool {
    k := make(map[int]int)
    for _, v := range nums{
        _, ok := k[v]
        if ok{
            return true
        }else{
            k[v] = 0
        }
    }
    return false
}
