func distributeCandies(candyType []int) int {
    candyMap := make(map[int]bool, len(candyType))
    for _, num := range candyType {
        candyMap[num] = true
    }
    if (len(candyType)/2) > len(candyMap) {
        return len(candyMap)
    } else {
        return len(candyType)/2
    }
    return 1
}
