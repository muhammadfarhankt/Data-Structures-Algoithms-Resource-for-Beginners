func canMakeSquare(grid [][]byte) bool {
    for i := 0; i < 2; i++ {
        black, white := 0, 0
        for j := 0; j < 2; j++ {
            if grid[i][j] == 'W' {
                white++
            } else {
                black++
            }
            if grid[i+1][j] == 'W' {
                white++
            } else {
                black++
            }
        }
        if black >= 3 || white >= 3 {
            return true
        }
        black, white = 0, 0
        for j := 1; j < 3; j++ {
            if grid[i][j] == 'W' {
                white++
            } else {
                black++
            }
            if grid[i+1][j] == 'W' {
                white++
            } else {
                black++
            }
        }
        if black >= 3 || white >= 3 {
            return true
        }
    }
    return false
}