func getEncryptedString(s string, k int) string {
    n := len(s)
    result := make([]byte, n)
    for i := 0; i < n; i++ {
        idx := (i + k) % n
        result[i] = s[idx]
    }
    return string(result)
}