func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    output := make([][]int, 0, rowIndex+1)
    output = append(output, []int{1})
    // fmt.Println(output)
    for i := 1; i <= rowIndex; i++ {
        newRow := make([]int, 0, i+1)
        newRow = append(newRow, 1)
        for j := 1; j < i; j++ {
            newRow = append(newRow, output[i-1][j-1]+output[i-1][j])
        }
        newRow = append(newRow, 1)
        output = append(output, newRow)
    }
    // fmt.Println(output)
    return output[rowIndex]
}