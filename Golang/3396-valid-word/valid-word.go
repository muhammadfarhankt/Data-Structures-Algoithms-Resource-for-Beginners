func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
    vowelMap := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
    vowel, consonant := 0, 0
    for _, char := range word {
        if vowelMap[char] {
            vowel++
        } else if char >= 'a' && char <= 'z' {
            consonant++
        } else if char >= '0' && char <= '9' {
            continue
        } else if char >= 'A' && char <= 'Z' {
            consonant++
        } else {
            return false
        }
    }
    if vowel >= 1 && consonant >= 1 {
            return true
    }
    return false
}