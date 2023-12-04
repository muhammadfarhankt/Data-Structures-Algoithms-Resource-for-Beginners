func checkString(s string) bool {
    for i, ch := range s {
        if ch == 'b' {
            if i == len(s)-1 {
                break
            }else if s[i+1] == 'a' {
                return false
            }
        } 
    }
    return true
}
