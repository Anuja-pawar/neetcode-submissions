func hasDuplicate(nums []int) bool {
    n := len(nums)
    recMP := map[int]struct{}{}
    
    for i:=0 ; i<n; i++{
        currVal := nums[i]
        if _, ok := recMP[currVal]; ok{
            return true
        }else{
            recMP[currVal] = struct{}{}
        }
    }
    return false
}
