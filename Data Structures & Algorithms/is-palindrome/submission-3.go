func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l<r {
		for l<r && !((s[l] >= 'a' && s[l] <= 'z') || (s[l] >= 'A' && s[l] <= 'Z') || (s[l] >= '0' && s[l] <= '9')){
			l++
		}
		for l<r && !((s[r] >= 'a' && s[r] <= 'z') || (s[r] >= 'A' && s[r] <= 'Z') || (s[r] >= '0' && s[r] <= '9')){
			r--
		}
		if strings.EqualFold(string(s[l]), string(s[r])){
			l++
			r--
		}else{

			return false
		}
	}
	return true
}

