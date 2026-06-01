func maxArea(heights []int) int {
    l := 0
    r := len(heights) - 1
    maxA := 0
    area := 0
    for l<r{
        if heights[l] < heights[r]{
            area = (r - l) * heights[l]
            l++
        }else{
            area = (r - l) * heights[r]
            r--
        }
        if area > maxA{
            maxA = area
        }
    }
    return maxA
}
