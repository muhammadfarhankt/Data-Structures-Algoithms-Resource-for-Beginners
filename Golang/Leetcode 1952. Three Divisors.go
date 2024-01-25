func isThree(n int) bool {
    if n <= 3 {
        return false
    }
    count := 0
    for i:=1; i<=n; i++ {
        if n % i == 0 {
            count++
        }
        if count > 3 {
            return false
        }
    }
    return count == 3
}
