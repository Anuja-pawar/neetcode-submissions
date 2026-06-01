func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    res := [][]int{}

    for i:=0; i<n; i++{
        curr := nums[i]
        if curr > 0{
            break
        }
        if i>0 && curr == nums[i-1]{
            continue
        }
        l := i+1
        r := n-1
        for l < r{
            sum := curr + nums[l] + nums[r]
            if sum < 0{
                l++
            }else if sum > 0{
                r--
            }else{
                res = append(res, []int{curr, nums[l], nums[r]})
                l++
                r--
                for l<r && nums[l] == nums[l-1]{
                    l++
                }
            }

        }
    }
    return res
}
