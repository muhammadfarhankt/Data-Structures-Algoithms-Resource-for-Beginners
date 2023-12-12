func findWordsContaining(words []string, x byte) []int {
    indexArray := []int{}
    for i, str := range words {
        for _, char := range str {
            if byte(char) == x {
                indexArray = append(indexArray, i)
                break
            }
        }
    }
    return indexArray
}
