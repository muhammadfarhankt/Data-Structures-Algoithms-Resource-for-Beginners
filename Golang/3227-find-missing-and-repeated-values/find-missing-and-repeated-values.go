func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    a, b := 0, 0
    gridMap := make(map[int]bool, n*n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if gridMap[grid[i][j]] == true {
                a = grid[i][j]
            } else {
                gridMap[grid[i][j]] = true
            }
        }
    }
    for i := 1; i <= n*n; i++ {
        if gridMap[i] != true {
            b = i
            break
        }
    }
    return []int{a, b}
}