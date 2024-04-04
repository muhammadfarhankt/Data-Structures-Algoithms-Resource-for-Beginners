func maxDepth(s string) int {
    depth, maxdepth := 0, 0
    for i:=0; i<len(s); i++ {
        if s[i] == '(' {
            depth++
        } else if s[i] == ')' {
            depth--
        }
        if depth > maxdepth {
            maxdepth = depth
        }
    }
    return maxdepth
}
