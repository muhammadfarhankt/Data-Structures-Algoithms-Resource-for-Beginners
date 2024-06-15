func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
    vowelMap := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
    vowel, consonant := false, false
    for i := 0; i < len(word); i++ {
        if vowelMap[word[i]] {
            vowel = true
        } else if word[i] >= 'a' && word[i] <= 'z' || word[i] >= 'A' && word[i] <= 'Z' {
            consonant = true
        } else if word[i] >= '0' && word[i] <= '9' {
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