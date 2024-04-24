func tribonacci(n int) int {
    if n == 0 {
        return 0
    } else if n < 3 { 
        return 1
    }
    arr := [38]int{0,1,1,2,4}
    for i := 3; i <= n; i++ {
        arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
    }
    return arr[n]
}