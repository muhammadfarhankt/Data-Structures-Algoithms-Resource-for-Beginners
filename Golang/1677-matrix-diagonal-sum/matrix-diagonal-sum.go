func diagonalSum(mat [][]int) int {
    size := len(mat)
    sum := 0
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if (i == j) || (i == size - j - 1) {
                sum += mat[i][j]
                fmt.Println(mat[i][j])
            }
        }
    }
    return sum
}