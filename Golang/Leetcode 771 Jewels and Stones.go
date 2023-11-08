func numJewelsInStones(jewels string, stones string) int {
    jewelMap := make(map[rune]int)
    count := 0
    for _, ch := range jewels {
        jewelMap[ch]++
    }
    for _, ch := range stones {
        if _, found := jewelMap[ch]; found {
            count++
        }
    }
    return count
}
