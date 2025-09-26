func maxFreqSum(s string) int {
    vowel, consonant := 0, 0
    lettersMap := make(map[byte]int, 26)
    for i := 0; i < len(s); i++ {
        lettersMap[s[i]]++
    }
    for char, val := range lettersMap {
        if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
            if val > vowel {
                vowel = val
            }
        } else {
            if val > consonant {
                consonant = val
            }
        }
    }
    return vowel + consonant
}