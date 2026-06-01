func isPalindrome(s string) bool {
    var strSlc []rune
    
    for _, i := range s{
       if unicode.IsLetter(i) || unicode.IsNumber(i){
            strSlc = append(strSlc, unicode.ToLower(i))
        }
    }
    n := len(strSlc)
    for i, j := 0, n-1; i < j ; i, j = i+1, j-1{
        if strSlc[i] != strSlc[j]{
            return false
        }
    }
    return true
}

