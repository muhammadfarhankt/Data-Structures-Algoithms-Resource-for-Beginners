func winningPlayerCount(n int, pick [][]int) int {
    playerColors := make([][]int, n)
    for i := 0; i < n; i++ {
        playerColors[i] = make([]int, 11)
    }
    for _, player := range pick {
        playerColors[player[0]][player[1]]++
    }
    count := 0
    for i := 0; i < n; i++ {
        for _, val := range playerColors[i] {
            if val > i {
                count++
                break
            }
        }
    }
    return count
}