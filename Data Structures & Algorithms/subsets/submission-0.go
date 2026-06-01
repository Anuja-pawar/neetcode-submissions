func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    subset := make([]int, 0) 
    var dfs func(i int)
    dfs = func (i int){
        fmt.Println("i ", i)
        if i >= len(nums){
            temp := make([]int, len(subset))
            copy(temp, subset) 
            res = append(res, temp)
            fmt.Println("res ", res)
            return
        }
        
        subset = append(subset, nums[i])
        fmt.Println("subset++ ",subset) 
        dfs(i + 1)
        
        subset = subset[:len(subset)-1]
        fmt.Println("subset-- ",subset) 
        dfs(i + 1)
    }
    dfs(0)
    return res
}
