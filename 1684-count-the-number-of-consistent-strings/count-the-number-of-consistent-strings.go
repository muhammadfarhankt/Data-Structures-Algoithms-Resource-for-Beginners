func countConsistentStrings(allowed string, words []string) int {
    allowedArr := make([]bool, 26)
    var count int
    for _, char := range allowed {
        allowedArr[char - 'a'] = true
    }
    for _, word := range words {
        var flag bool
        for _, char := range word {
            if !allowedArr[char - 'a'] {
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