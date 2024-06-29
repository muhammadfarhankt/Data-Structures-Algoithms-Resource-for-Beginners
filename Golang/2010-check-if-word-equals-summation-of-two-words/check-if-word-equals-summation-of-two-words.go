func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    word1, word2, targetSum := 0, 0, 0
    for _, char := range firstWord {
        word1 = word1 * 10 + int(char - 'a')
    }
    for _, char := range secondWord {
        word2 = word2 * 10 + int(char - 'a')
    }
    for _, char := range targetWord {
        targetSum = targetSum * 10 + int(char - 'a')
    }
    // fmt.Println(word2, targetSum)
    if word1 + word2 == targetSum {
        return true
    }
    return false
}