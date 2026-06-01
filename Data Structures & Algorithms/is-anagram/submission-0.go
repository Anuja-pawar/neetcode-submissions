func isAnagram(s string, t string) bool {
    n := len(s)
    m := len(t)
    if(n != m){
        return false
    }
    sMap := make(map[string]int)
    tMap := make(map[string]int)
    for _, v := range s{
        sMap[string(v)] = sMap[string(v)] + 1   
    }
    for _, v := range t{
        tMap[string(v)] = tMap[string(v)] + 1   
    }
    for k, v := range sMap{
        if v == tMap[k]{
            continue
        }else{
            return false
        }
    }
    return true
}
