func maxLengthBetweenEqualCharacters(s string) int {
    maxLength := -1
    arrMap := make(map[rune]int)
    for idx, char := range s {
        if j, ok := arrMap[char]; ok {
            if (idx - j - 1) > maxLength {
                maxLength = idx - j - 1
            }
        } else {
            arrMap[char] = idx
        }
    }
    return maxLength
}
