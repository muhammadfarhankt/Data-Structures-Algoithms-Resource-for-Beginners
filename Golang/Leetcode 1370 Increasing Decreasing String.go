func sortString(s string) string {
    charMap := make(map[rune]int)
    outputString := make([]rune, 0, len(s))
    for _, char := range s {
        charMap[char]++
    }
    for len(outputString) < len(s) {
        for char := 'a'; char <= 'z'; char++ {
            if charMap[char] > 0 {
                outputString = append(outputString, char)
                charMap[char]--
            }
        }
        for char := 'z'; char >= 'a'; char-- {
            if charMap[char] > 0 {
                outputString = append(outputString, char)
                charMap[char]--
            }
        }
    }
    return string(outputString)
}
