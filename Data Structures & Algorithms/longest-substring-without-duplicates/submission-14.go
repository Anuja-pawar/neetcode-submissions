func lengthOfLongestSubstring(s string) int {
    l := 0
    r := 0
    longest := 0
    seen := make(map[byte]int)
    for r < len(s){
        if idx, found := seen[s[r]]; found{
            l = max(idx+1, l)
        }
        seen[s[r]] = r
        if r - l + 1 > longest{
            longest = r - l + 1
        }
        r++
    }
    return longest
}
