func countGoodSubstrings(s string) int {
    count := 0
    if len(s) < 3 {
        return 0
    }
    stringSet := make(map[byte]int)
    for i := 0; i < 3; i++ {
        stringSet[s[i]]++
    }
    if len(stringSet) == 3 {
        count++
    }
    for i := 3; i < len(s); i++ {
        if stringSet[s[i-3]] == 1 {
            delete(stringSet, s[i-3])
        } else {
            stringSet[s[i-3]]--
        }
        stringSet[s[i]]++
        if len(stringSet) == 3 {
            count++
        }
    }
    return count
}