func commonFactors(a int, b int) int {
    small, count := a, 0
    if a > b {
        small = b
    }
    for i := 1; i <= small; i++ {
        if a % i == 0 && b % i == 0 {
            count++
        }
    }
    return count
}