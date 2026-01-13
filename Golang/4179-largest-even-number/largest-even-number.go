func largestEven(s string) string {
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == '2' {
            return s[:i+1]
        }
    }
    return ""
}