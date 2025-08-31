func getLeastFrequentDigit(n int) int {
    arr := make([]int, 10)
    for n > 0 {
        temp := n % 10
        n = n / 10
        arr[temp]++
    }
    fmt.Println(arr)
    for i := 1; i <= 9; i++ {
        for j := 0; j <= 9; j++ {
            if arr[j] == i {
                return j
            }
        }
    }
    return 0
}