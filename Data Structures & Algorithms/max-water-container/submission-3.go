func maxArea(heights []int) int {
    l := 0 
    r := len(heights)-1
    resArea := 0
    for l < r{
        height := min(heights[l], heights[r])
        currArea := height * (r - l)
        if currArea > resArea {
            resArea = currArea
        }
        if heights[l] < heights[r]{
            l++
        }else{
            r--
        }
    }
    return resArea

}
