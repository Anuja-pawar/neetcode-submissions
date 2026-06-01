func lengthOfLongestSubstring(s string) int {
    l := 0
    r := 0
    longest := 0
    seen := make(map[string]int)
    for r < len(s){
        if seen[string(s[r])] > 0{
            for l < r && seen[string(s[r])] > 0{
                seen[string(s[l])]--
                l++
            }
        }
        seen[string(s[r])]++
        curr := r - l + 1
        longest = max(curr, longest)
        r++
    }
    return longest
}
