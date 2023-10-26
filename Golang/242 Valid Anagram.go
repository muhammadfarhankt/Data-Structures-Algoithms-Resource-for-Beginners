func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    if s == t {
        return true
    }
    charMap := make(map[rune]int)
    for _,ch := range s {
        charMap[ch]++
    }
    for _,ch := range t {
        charMap[ch]--
        if (charMap[ch] < 0 ) {
            return false
        }
    }
    return true
}
