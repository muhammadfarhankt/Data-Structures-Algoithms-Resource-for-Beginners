func replaceElements(arr []int) []int {
    output := make([]int, len(arr))
    for i := 0; i < len(arr); i++ {
        max := -1
        for j := i+1; j < len(arr); j++ {
            if arr[j] > max {
                max = arr[j]
            }
        }
        output[i] = max
    }
    return output
}