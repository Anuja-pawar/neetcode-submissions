func minWindow(s string, t string) string {
    if len(s) < len(t){
		return ""
	}
	tMap := make(map[byte]int)

	for i:=0; i<len(t); i++{
		tMap[t[i]]++
	}
	currMap := make(map[byte]int)
	have, need := 0, len(tMap)
	l := 0
	res, resLen := []int{-1,-1}, math.MaxInt32

	for r:=0; r<len(s); r++{
		ch := s[r]
		currMap[ch]++
		if tMap[ch] > 0 && tMap[ch] == currMap[ch]{
			have++
		}

		for have == need{
			if (r-l+1) < resLen{
				res = []int{l,r}
				resLen = r - l + 1
			}
			currMap[s[l]]--
			if tMap[s[l]] > 0 && tMap[s[l]] > currMap[s[l]]{
				have--
			}
			l++
		}
	}
	if res[0] == -1{
		return ""
	}
	return s[res[0]:res[1]+1]
}
