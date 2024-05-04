func mostWordsFound(sentences []string) int {
    max := 0
    for _, sentence := range sentences {
        space := 0
        for _, char := range sentence {
            if char == ' ' {
                space++
            }
        }
        if space+1 > max {
            max = space+1
        }
    }
    return max
}