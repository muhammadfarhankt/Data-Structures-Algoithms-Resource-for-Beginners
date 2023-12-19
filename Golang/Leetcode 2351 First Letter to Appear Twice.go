func repeatedCharacter(s string) byte {
    charMap := make(map[rune]int)
    for _, ch := range s {
        charMap[ch]++
        if charMap[ch] == 2 {
            return byte(ch)
        }
    }
    return ' '
}
