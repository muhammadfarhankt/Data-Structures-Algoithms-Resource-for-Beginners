func maxDepth(s string) int {
    max, depth := 0, 0
    for _, ch := range s {
        if ch == '(' {
            fmt.Println("(")
            depth++
            fmt.Println("depth : ", depth)
            if depth > max {
                max = depth
                fmt.Println("max : ", max)
            }
        } else if ch == ')' {
            fmt.Println(")")
            depth--
        }
    }
    return max
}