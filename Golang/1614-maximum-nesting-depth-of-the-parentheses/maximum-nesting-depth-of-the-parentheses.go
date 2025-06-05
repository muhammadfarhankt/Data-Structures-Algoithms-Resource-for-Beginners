func maxDepth(s string) int {
    max, depth := 0, 0
    for _, ch := range s {
        if ch == '(' {
            depth++
            if depth > max {
                max = depth
            }
        } else if ch == ')' {
            depth--
        }
    }
    return max
}