func construct2DArray(original []int, m int, n int) [][]int {
    if n * m != len(original) {
        return [][]int{}
    }

    output := make([][]int, m)
    
    for i, idx := 0, 0; i < m; i++ {
        output[i] = make([]int, n)
        for j := 0; j < n; j++ {
            output[i][j] = original[idx]
            idx++
        }
    }
    
    return output
}