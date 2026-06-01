func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)

    for i:=0 ; i<len(nums); i++{
        diff := target - nums[i]
        if _, ok:= mp[diff]; ok{
            return []int{mp[diff], i}
        }else{
            val := nums[i]
            mp[val] = i
        }
        fmt.Println(mp)
    }
    return []int{}
}
