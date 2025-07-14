func truncateSentence(s string, k int) string {
    for i, space := 0, 0; i < len(s) ; i++ {
        if s[i] == ' ' {
            space++
        }
        if space == k {
            return s[:i]
        }
    }
    return s
}