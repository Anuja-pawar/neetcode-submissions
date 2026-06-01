func isValid(s string) bool {
    stack := make([]rune, 0)
    closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if open, exists := closeToOpen[c]; exists {
            if len(stack) != 0 {
                lastInd := len(stack)-1
                top := stack[lastInd]
                stack = stack[:lastInd]
                if  top != open {
                    return false
                }
            } else {
                return false
            }
        } else {
            stack = append(stack, c)
        }
    }
    if len(stack) != 0{
        return false
    }
    return true
}