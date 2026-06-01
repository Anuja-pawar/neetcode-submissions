func groupAnagrams(strs []string) [][]string {
    res := make(map[string][]string)
    for _, s := range strs{
        sortedS := SortThis(s)
        res[sortedS] = append(res[sortedS], s) 
    }
    var result [][]string
    for _, grp := range res{
        result = append(result, grp)
    }
    return result
}
func SortThis(s string) string{
        characters := []rune(s)
        sort.Slice(characters, func(i, j int) bool{
            return characters[i] < characters[j]
        })
        return string(characters)
}
 
