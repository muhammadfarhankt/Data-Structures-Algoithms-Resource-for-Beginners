func numberOfSpecialChars(word string) int {
    small, capital := make(map[rune]int, 26), make(map[rune]int, 26)
    count := 0
    for i, char := range word {
        if char >= 'a' && char <= 'z' {
            small[char] = i
        } else if char >= 'A' && char <= 'Z' {
            if _, exists := capital[char]; !exists {
                capital[char] = i
            }
        }
    }
    for key, value := range small {
        if value < capital[key-32] {
            count++
        }
    }
    fmt.Println(small)
    fmt.Println(capital)
    return count
}