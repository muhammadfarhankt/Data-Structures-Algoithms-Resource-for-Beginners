func mergeAlternately(word1 string, word2 string) string {
    l1, l2 := len(word1), len(word2)
    result := make([]byte, 0, (l1+l2))
    for i, j := 0, 0; i < l1 || j < l2; {
        if i < l1 {
            result = append(result, word1[i])
            i++
        }
        if j < l2 {
            result = append(result, word2[j])
            j++
        }
    }
    return string(result)
}