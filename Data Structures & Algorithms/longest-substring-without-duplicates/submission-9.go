func lengthOfLongestSubstring(s string) int {
    if len(s) == 0{
        return 0
    }else if len(s)==1{
        return 1
    }
    hashMap := make(map[rune]struct{})
    res := 1
    curr_res := 0
    l, r := 0, 1
    for l < r && r < len(s){
        lVal := rune(s[l])
        hashMap[lVal]=struct{}{}
        rVal := rune(s[r])
        if _,ok:=hashMap[rVal];ok{
            hashMap = make(map[rune]struct{})
            l = l+1
            r = l+1
            if curr_res > res{
                res = curr_res
            }
        }else{
            hashMap[rVal]=struct{}{}
            r++
            curr_res = len(hashMap)
        }

    }
    if curr_res > res{
        res = curr_res
    }
    return res
}
            
