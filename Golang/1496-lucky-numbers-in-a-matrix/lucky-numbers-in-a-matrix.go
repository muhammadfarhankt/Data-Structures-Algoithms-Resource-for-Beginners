func luckyNumbers (matrix [][]int) []int {
    output := make([]int, 0)
    minSet := make(map[int]bool)
    m, n := len(matrix), len(matrix[0])
    for i := 0; i < m; i++ {
        min := matrix[i][0]
        for j := 1; j < n; j++ {
            if matrix[i][j] < min {
                min = matrix[i][j]
            }
        }
        minSet[min] = true
    }
    for i := 0; i < n; i++ {
        max := matrix[0][i]
        for j := 1; j < m; j++ {
            if matrix[j][i] > max {
                max = matrix[j][i]
            }
        }
        if _, exists := minSet[max]; exists {
            output = append(output, max)
        }
    }
    return output
}