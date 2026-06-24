func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1

    for l < r{
        for l < r && !isAlphaNum(rune(s[l])) {
            l++
        }
        for l < r && !isAlphaNum(rune(s[r])) {
            r--
        }
        if strings.ToLower(string(s[l])) == strings.ToLower(string(s[r])) {
            l++
            r--
        }else{
            return false
        }

    }
    return true

}

func isAlphaNum(i rune) bool{
	return unicode.IsLetter(i) || unicode.IsDigit(i)
}
