func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    a, b := 0, 0
    gridArr := make([]bool, n*n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if gridArr[grid[i][j]-1] == true {
                a = grid[i][j]
            } else {
                gridArr[grid[i][j]-1] = true
            }
        }
    }
    // fmt.Println(gridArr)
    for i := 1; i <= n*n; i++ {
        if gridArr[i-1] != true {
            b = i
            break
        }
    }
    return []int{a, b}
}