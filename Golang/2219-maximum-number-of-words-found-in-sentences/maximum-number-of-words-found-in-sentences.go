func mostWordsFound(sentences []string) int {
    maxCount := 0
    for i := 0; i < len(sentences); i++ {
        count := 1
        for j := 0; j < len(sentences[i]); j++ {
            if sentences[i][j] == ' ' {
                count += 1
            }
        }
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}