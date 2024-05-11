func satisfiesConditions(grid [][]int) bool {
    column, row := len(grid[0]), len(grid)
    for i := 0; i < row-1; i++ {
        for j := 0; j < column; j++ {
            if grid[i][j] != grid[i+1][j] {
                return false
            }
        }
    }
    for i := 0; i < row; i++ {
        for j := 0; j < column-1; j++ {
            if grid[i][j] == grid[i][j+1] {
                return false
            }
        }
    }
    return true
}