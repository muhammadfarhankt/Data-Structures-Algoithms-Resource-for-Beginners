func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
    vowelMap := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
    vowel, consonant := false, false
    for _, char := range word {
        if vowelMap[char] {
            vowel = true
        } else if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
            consonant = true
        } else if char >= '0' && char <= '9' {
            continue
        } else {
            return false
        }
    }
    if vowel && consonant {
        return true
    }
    return false
}