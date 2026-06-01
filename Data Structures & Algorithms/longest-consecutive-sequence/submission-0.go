func longestConsecutive(nums []int) int {
    myMap := make(map[int]int)
    res := 0

    for _, num := range nums{
        if myMap[num] == 0{
            left := myMap[num-1]
            right := myMap[num+1]
            sum := left + right + 1
            myMap[num] = sum
            myMap[num - left] = sum
            myMap[num + right] = sum
            if sum > res{
                res = sum
            }
        }
    }
    return res
}