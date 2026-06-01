func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int)

    for _, n := range nums{
        freqMap[n]++ 
    }
    temp := make([][]int, 0)
    for num, count := range freqMap{
        temp = append(temp, []int{count, num})
    }

    sort.Slice(temp, func (i, j int) bool{
        return temp[i][0] > temp[j][0]
    })
    res := make([]int, 0)
    for _, n := range temp{
        if len(res) < k{
            res = append(res, n[1])
        }else{
            break
        }   
    }
    return res
}
