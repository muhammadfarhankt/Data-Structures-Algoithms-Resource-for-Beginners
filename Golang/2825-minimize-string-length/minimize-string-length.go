func minimizedStringLength(s string) int {
    charSet := make(map[rune]bool)
    for _, char := range s {
        charSet[char] = true
    }
    return len(charSet)
}