func trap(height []int) int {
    if len(height) <= 0{
        return 0
    }
    l, r := 0, len(height) -1
    maxLeft := height[l]
    maxRight := height[r]
    tw := 0
    res := 0
    for l <= r{
        if maxLeft < height[l]{
            maxLeft = height[l]
        }
        if maxRight < height[r]{
            maxRight = height[r]
        }
        if maxLeft <= maxRight{
            tw = maxLeft - height[l]
            l++
        }else{
            tw = maxRight- height[r]
            r--
        }
        if tw > 0{
            res = res + tw
        }
    }
    return res
}
