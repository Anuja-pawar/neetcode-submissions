func characterReplacement(s string, k int) int {
    res := 0
    charMap := make(map[byte]bool)

    for i := 0; i < len(s); i++ {
        charMap[s[i]] = true
    }

    for ch := range charMap{
        l := 0
        count := 0

        for r:=0 ; r<len(s); r++{
            if s[r] == ch{
                count++
            }

            for (r - l + 1) - count > k{
                if s[l] == ch{
                    count--
                }
                l++
            }
            res = max(res, r-l+1)
        }
    }
    return res
}