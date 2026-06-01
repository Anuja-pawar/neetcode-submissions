func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    total := 0
    l, r := 0, len(height) -1
    maxL, maxR := height[l], height[r]
    for l < r{
        if maxL <= maxR{
            l++
            maxL = max(maxL, height[l])
            total += maxL - height[l]
        }else{
            r--
            maxR = max(maxR, height[r])
            total += maxR - height[r]
        }
    }
    return total
}
