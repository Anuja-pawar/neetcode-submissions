func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1
    for l < r{
        for !isAlphaNum(rune(s[l])) && l < r{
            l++
        }
        for !isAlphaNum(rune(s[r])) && l < r{
            r--
        }
        if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])){
            return false
        }
        l++
        r--
    }
    return true

}

func isAlphaNum(i rune) bool{
	return unicode.IsLetter(i) || unicode.IsDigit(i)
}
