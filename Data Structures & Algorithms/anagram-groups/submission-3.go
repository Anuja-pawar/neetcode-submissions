func groupAnagrams(strs []string) [][]string {
    charCount := make(map[[26]int][]string)
    for _, word := range strs{
        var chArr [26]int
        for _, ch := range word{
            chArr[ch - 'a']++
        }
        charCount[chArr] = append(charCount[chArr], word)
    }

    res := make([][]string, 0)
    for _, val := range charCount{
        res = append(res, val)
    }
    return res
}
