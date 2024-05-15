func findPeaks(mountain []int) []int {
    size := len(mountain)
    result := make([]int, 0, size-2)
    for i := 1; i < size-1; i++ {
        if (mountain[i] > mountain[i - 1]) && (mountain[i] > mountain[i + 1]) {
            result = append(result, i)
        }
    }
    return result
}