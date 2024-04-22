func numberOfSpecialChars(word string) int {
    small, capital := [26]int{}, [26]int{}
    count := 0
    for i, char := range word {
        if char >= 'a' && char <= 'z' {
            small[int(char - 'a')] = i+1
        } else {
            if capital[int(char - 'A')] == 0 {
                capital[int(char - 'A')] = i+1
            }
        }
    }
    for i := 0; i < 26; i++ {
        if small[i] != 0 && capital[i] > small[i] {
            count++
        }
    }
    return count
}