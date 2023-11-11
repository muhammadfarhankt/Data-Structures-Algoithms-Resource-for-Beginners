func areOccurrencesEqual(s string) bool {
    stringMap := make(map[rune]int)
    for _, ch := range s {
        stringMap[ch]++
    }
    value := stringMap[rune(s[0])]
    for _, count := range stringMap {
        if value != count {
            return false
        }
        value = count
    }
    return true
}
