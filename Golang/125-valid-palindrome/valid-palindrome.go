func isPalindrome(s string) bool {
    size := len(s)
    str := make([]byte, 0, size)
    for i := 0; i < size; i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            str = append(str, s[i])
        } else if s[i] >= 'A' && s[i] <= 'Z' {
            str = append(str, byte(s[i] + 32))
        } else if s[i] >= '0' && s[i] <= '9' {
            str = append(str, s[i])
        }
    }
    if len(str) == 0 {
        return true
    }
    for i := 0; i < len(str)/2; i++ {
        if str[i] != str[len(str)-i-1] {
            return false
        }
    }
    return true
}