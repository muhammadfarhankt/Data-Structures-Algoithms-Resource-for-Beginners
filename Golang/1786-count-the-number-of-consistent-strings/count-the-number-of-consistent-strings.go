func countConsistentStrings(allowed string, words []string) int {
    allowedMap := make(map[rune]bool)
    var count int
    for _, char := range allowed {
        allowedMap[char] = true
    }
    for _, word := range words {
        var flag bool
        for _, char := range word {
            if _, found := allowedMap[char]; !found {
                flag = true
                break
            }
        }
        if !flag {
            count++
        }
    }
    return count
}