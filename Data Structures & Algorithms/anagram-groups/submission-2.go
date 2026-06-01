func groupAnagrams(strs []string) [][]string {
    charCount := make(map[[26]int][]string)

    for _, str := range(strs){
        var arr [26]int
        for _, ch := range(str) {
            arr[ch - 'a']++
        }
        charCount[arr] = append(charCount[arr], str) 
    }
    vals := make([][]string, 0)

    for _, v := range charCount {
        vals = append(vals, v)
    }
    return vals
}
