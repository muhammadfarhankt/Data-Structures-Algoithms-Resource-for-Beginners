func replaceElements(arr []int) []int {
    for i, max := len(arr)-1, -1; i >= 0; i-- {
        if arr[i] > max {
            arr[i], max = max, arr[i]
        } else {
            arr[i] = max
        }
    }
    return arr
}