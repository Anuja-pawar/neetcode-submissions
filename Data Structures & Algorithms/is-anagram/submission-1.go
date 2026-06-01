func isAnagram(s string, t string) bool {
    sArr := []rune(s)
    tArr := []rune(t)
    sCharMp := make(map[string]int)
    tCharMp := make(map[string]int)
    if len(s) != len(t){
        return false
    }

    for i:=0; i<len(sArr); i++{
        a := string(sArr[i])
        b := string(tArr[i])
        sCharMp[a]++
        tCharMp[b]++
    }

    for k := range sCharMp{
        if sCharMp[k] != tCharMp[k]{
            return false
        }
    }
    return true

}
