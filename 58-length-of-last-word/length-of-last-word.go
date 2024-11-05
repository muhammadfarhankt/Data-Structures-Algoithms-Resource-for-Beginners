func lengthOfLastWord(s string) int {
    current, lastLength := 0, 0
    for _, char := range s {
        if char != ' ' {
            current++
        } else {
            current = 0
        }
        if current > 0 {
            lastLength = current
        }
    }
    return lastLength
}