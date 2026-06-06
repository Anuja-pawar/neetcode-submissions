func checkInclusion(s1 string, s2 string) bool{
	a := len(s1)
	b := len(s2)

	if a > b {
		return false
	}

	var s1Count, s2Count [26] int

	for i:=0; i<a; i++{
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	if s1Count == s2Count{
		return true
	}

	for i:=a; i<b; i++{
		s2Count[s2[i]-'a']++
		s2Count[s2[i-a]-'a']--
		if s1Count == s2Count{
			return true
		}
	}
	return false

}