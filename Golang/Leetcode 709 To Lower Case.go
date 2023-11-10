func toLowerCase(s string) string {
    chars := ""
    for i,ch := range s {
        if ch >= 65 && ch <= 90 {
            chars += string(ch + 32)
        } else {
            chars += string(s[i])
        }
    }
    return chars
}
