func topKFrequent(nums []int, k int) []int {
    cnt := make(map[int]int)
    for _, n := range nums{
        cnt[n]++
    }

    arr := make([][2]int, 0, len(cnt))
    for num, count := range cnt{
        arr = append(arr, [2]int{count, num})
    }
    sort.Slice(arr, func(i,j int) bool{
        return arr[i][0] > arr[j][0]
    })

    var res []int
    for i := 0; i <k; i++{
        res = append(res, arr[i][1])
    }
    return res

}
