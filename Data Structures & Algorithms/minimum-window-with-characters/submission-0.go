func minWindow(s string, t string) string {
    if len(s) < len(t){
        return ""
    }

    targetMap := make(map[rune]int)
    for _, ch := range t {
        targetMap[ch]++
    }

    currWindowMap := make(map[rune]int)
    have, need := 0, len(targetMap)
    l := 0
    res, resLength := []int{-1,-1}, math.MaxInt32

    for r := 0; r < len(s); r++{
        ch := rune(s[r])
        currWindowMap[ch]++
        if targetMap[ch] > 0 && targetMap[ch] == currWindowMap[ch] {
            have++
        }

        for have == need{
            if (r - l + 1) < resLength{
                res = []int{l,r}
                resLength = r - l + 1
            }
            currWindowMap[rune(s[l])]--
            if targetMap[rune(s[l])] > 0 && targetMap[rune(s[l])] > currWindowMap[rune(s[l])]{
                have--
            }
            l++    
        }
    }
    if res[0] == -1{
        return ""
    }

    return s[res[0]:res[1]+1]
    
}
