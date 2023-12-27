func isAcronym(words []string, s string) bool {
    if len(s) != len(words) {
        return false
    }
    for i:=0; i<len(s); i++ {
        if s[i] != words[i][0] {
            return false
        }
    }
    return true
}
