func commonFactors(a int, b int) int {
    count, low := 0, 0
    if a < b {
        low = a
    } else {
        low = b
    }
    for i := 1; i <= low; i++ {
        if a % i == 0 && b % i == 0 {
            count++
        }   
    }
    return count
}