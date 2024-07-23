func replaceElements(arr []int) []int {
    for i, max := len(arr)-1, -1; i >= 0; i-- {
        temp := max
        if arr[i] > max {
            max = arr[i]
        }
        arr[i] = temp
    }
    return arr
}