func truncateSentence(s string, k int) string {
    count := 0
    for i, val := range s {
        if val == ' ' {
            count++
        }
        if count == k {
            return s[0:i]
        }
    }
    return s
}
