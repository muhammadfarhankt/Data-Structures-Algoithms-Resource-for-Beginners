//Using maps
func isValid(s string) bool {
    sStack := []rune{}
    sMap := map[rune]rune {
        ')':'(',
        ']':'[',
        '}':'{',
    }
    for _, ch := range s {
        if ch == '{' || ch == '(' || ch == '[' {
            sStack = append(sStack, ch)
        } else if ch == '}' || ch == ')' || ch == ']' {
            if len(sStack) == 0 || sMap[ch] != sStack[len(sStack)-1] {
                return false
            }
            sStack = sStack[:len(sStack)-1]
        }
    }
    return len(sStack) == 0
}

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
