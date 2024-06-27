func vowelStrings(words []string, left int, right int) int {
    count := 0
    for i := left; i <= right; i++ {
        word := words[i]
        last := len(word) - 1
        if (word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u') && (word[last] == 'a' || word[last] == 'e' || word[last] == 'i' || word[last] == 'o' || word[last] == 'u') {
            count++
        }
    }
    return count
}