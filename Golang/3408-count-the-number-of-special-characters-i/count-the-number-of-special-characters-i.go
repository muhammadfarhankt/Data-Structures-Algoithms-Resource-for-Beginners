func numberOfSpecialChars(word string) int {
    small, capital := [26]int{}, [26]int{}
    count := 0
    for _, char := range word {
        if char >= 'a' && char <= 'z' {
            small[int(char - 'a')] = 1
        } else if char >= 'A' && char <= 'Z' {
            capital[int(char - 'A')] = 1
        }
    }
    for i := 0; i < 26; i++ {
        if small[i] == 1 && capital[i] == 1 {
            count++
        }
    }
    return count
}