func countGoodTriplets(arr []int, a int, b int, c int) int {
    size, count := len(arr), 0
    for i := 0; i < size; i++ {
        for j := i+1; j < size; j++ {
            for k := j + 1; k < size; k++ {
                if abs(arr[i] - arr[j]) <= a && abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k]) <= c {
                    count++
                }
            }
        }
    }
    return count
}
func abs(num int) int {
    if num >= 0{
        return num
    }
    return -num
}