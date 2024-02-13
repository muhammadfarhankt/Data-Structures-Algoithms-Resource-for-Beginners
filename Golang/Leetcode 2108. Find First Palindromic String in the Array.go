func firstPalindrome(words []string) string {
    for _, word := range words {
        if checkPalindrome(word) {
            return word
        }
    }
    return ""
}
func checkPalindrome(word string) bool {
    l := len(word)
    for i:=0; i<l/2; i++ {
        if word[i] != word[l-i-1] {
            return false
        }
    }
    return true
}
