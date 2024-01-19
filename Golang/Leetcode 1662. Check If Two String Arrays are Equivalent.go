func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    str1 := ""
    str2 := ""
    for _, str := range word1 {
        str1 += str
    }
    for _, str := range word2 {
        str2 += str
    }
    if str1 == str2 {
        return true
    }
    return false
}
