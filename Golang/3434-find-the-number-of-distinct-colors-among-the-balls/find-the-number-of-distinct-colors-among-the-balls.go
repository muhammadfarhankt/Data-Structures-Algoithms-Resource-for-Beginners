func queryResults(limit int, queries [][]int) []int {
    ballColors := make(map[int]int)
    colorCount := make(map[int]int)
    result := make([]int, len(queries))
    distinctColors := 0

    for i, query := range queries {
        ball := query[0]
        color := query[1]

        if prevColor, exists := ballColors[ball]; exists {
            colorCount[prevColor]--
            if colorCount[prevColor] == 0 {
                distinctColors--
            }
        }

        ballColors[ball] = color
        colorCount[color]++
        if colorCount[color] == 1 {
            distinctColors++
        }

        result[i] = distinctColors
    }

    return result
}