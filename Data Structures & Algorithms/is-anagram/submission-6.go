func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    var s_counter = make(map[rune]int, len(s))
    var t_counter = make(map[rune]int, len(t))
    for _, ch := range(s){
        s_counter[ch]++
    }
    for _, ch := range(t){
        t_counter[ch]++
    }

    for _, ch := range(s){
        if s_counter[ch] != t_counter[ch]{
            return false
        }
    }
    return true
}
