func checkIfPangram(sentence string) bool {
    // Create Hash Table for storing alphabets
    alphabets := make(map[rune]int)
    if len(sentence) < 26 {
        return false
    }
    // Iterate through each character
    for _, singleChar := range sentence {
        alphabets[singleChar]++
    }
    if len(alphabets) == 26 {
        return true
    } else {
        return false
    }
}
