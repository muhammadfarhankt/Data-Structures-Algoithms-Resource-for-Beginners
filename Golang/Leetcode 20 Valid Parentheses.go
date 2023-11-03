func isValid(s string) bool {
    // if len(s) % 2 != 0 {
    //     return false
    // }
    stack := []rune{}
    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack,ch)
        } else if ch == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' {
            stack = stack[:len(stack)-1]
        } else if ch == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
            stack = stack[:len(stack)-1]
        } else if ch == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
            stack = stack[:len(stack)-1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}
