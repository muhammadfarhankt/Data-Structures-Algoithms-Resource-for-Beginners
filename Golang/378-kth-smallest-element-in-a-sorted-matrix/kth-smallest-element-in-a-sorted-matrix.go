func kthSmallest(matrix [][]int, k int) int {
    size := len(matrix[0])
    arr := make([]int, 0, (size*size))
    for _, row := range matrix {
        for _, element := range row {
            arr = append(arr, element)
        }
    }
    sort.Ints(arr)
    return arr[k-1]
}