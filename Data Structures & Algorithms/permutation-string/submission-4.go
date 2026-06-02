func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2){
        return false
    }

	s1Count := make([]int, 26)
	s2Count := make([]int, 26)
    countMatch := 0

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

    for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			countMatch++
		}
	}

    l := 0
	for r := len(s1); r < len(s2); r++ {
		if countMatch == 26 {
			return true
		}
        index := s2[r] - 'a'
        s2Count[index]++
        if s1Count[index] == s2Count[index] {
			countMatch++
		} else if s1Count[index]+1 == s2Count[index] {
			countMatch--
		}
        index = s2[l] - 'a'
		s2Count[index]--
        if s1Count[index] == s2Count[index] {
			countMatch++
		} else if s1Count[index]-1 == s2Count[index] {
			countMatch--
		}
		l++
    }
    return countMatch == 26
}
