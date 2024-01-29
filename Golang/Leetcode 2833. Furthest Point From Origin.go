func furthestDistanceFromOrigin(moves string) int {
    movesMap := make(map[rune]int)
    for _, move := range moves {
        movesMap[move]++
    }
    if movesMap['R'] > movesMap['L'] {
        return movesMap['_'] + movesMap['R'] - movesMap['L']
    } else {
        return movesMap['_'] + movesMap['L'] - movesMap['R']
    }
    return 0
}
