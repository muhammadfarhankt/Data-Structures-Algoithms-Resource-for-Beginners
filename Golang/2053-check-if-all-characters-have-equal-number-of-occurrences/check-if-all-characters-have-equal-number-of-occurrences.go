func areOccurrencesEqual(s string) bool {
    charArr := [26]int{}
    value := 0
    for _, char := range s {
        charArr[char - 'a']++
        value = charArr[char - 'a']
    }
    for _, val := range charArr {
        if val != 0 && val != value {
            return false
        }
    }
    return true
}