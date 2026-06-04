func checkInclusion(s1 string, s2 string) bool{
	
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2{
		return false
	}
	var s1Count, s2Count [26]int

	for i:=0;i<n1;i++{
		s1Count[s1[i] - 'a']++
		s2Count[s2[i] - 'a']++
	}

	if s1Count == s2Count{
		return true
	}

	for i:=n1; i<n2; i++{
		s2Count[s2[i]-'a']++
		s2Count[s2[i-n1]-'a']--
		if s1Count == s2Count{
			return true
		}
	}
	return false
}