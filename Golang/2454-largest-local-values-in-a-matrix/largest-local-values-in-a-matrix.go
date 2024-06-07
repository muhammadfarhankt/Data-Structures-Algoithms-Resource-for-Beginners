func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    // fmt.Println(n)
    output := make([][]int, (n-2), (n-2))
    for i := 0; i < n-2; i++ {
        output[i] = make([]int, n-2)
    }
    // fmt.Println(output)
    for i := 0; i < n-2; i++ {
        for j := 0; j < n-2; j++ {
            max := grid[i][j]
            for k := i; k < i+3; k++ {
                for l := j; l < j+3; l++ {
                    if grid[k][l] > max {
                        max = grid[k][l]
                    }
                }
            }
            output[i][j] = max
        }

    }
    return output
}