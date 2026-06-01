func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    for _, val := range tokens{
        switch val {
            case "+":
                a := stack[len(stack)-1]
                b := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, b+a)
            case "-":
                a := stack[len(stack)-1]
                b := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, b-a)
            case "*":
                a := stack[len(stack)-1]
                b := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, b*a)
            case "/":
                a := stack[len(stack)-1]
                b := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, b/a)
            default:
                num, _ := strconv.Atoi(val)
                stack = append(stack, num)
        }
    }
    return stack[0]
}

