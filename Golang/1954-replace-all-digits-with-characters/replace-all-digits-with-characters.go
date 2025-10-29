func replaceDigits(s string) string {
    sByte := []byte(s)
    for i := 1; i < len(s); i += 2 {
        sByte[i] += sByte[i-1] - '0'
    }
    return string(sByte)
}