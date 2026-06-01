func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) <= 0{
        return []int{}
    }
    res := make([]int, 0, len(nums)-k+1)
    curr := make([]int, 0)

    for i:=0; i<len(nums); i++{
        // removing the element from current window
        //  if it's out of the fixed window size
        if len(curr) > 0 && curr[0] <= i - k{
            curr = curr[1:]
        }
        // if elements in the window is smaller than the curr element
        // that means they can't be max value, then remove it
        for len(curr) > 0 && nums[curr[len(curr)-1]] < nums[i]{
            curr = curr[:len(curr)-1]
        }
        // append current element
        curr = append(curr, i)
        // start adding result once we hit the first window
        if i >= k - 1{
            res = append(res, nums[curr[0]])
        }
    }
    return res
}