func climbStairs(n int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    }
    a, b, fib := 1, 1, 0
    for i := 3; i<=n; i++ {
        fib = a + b
        a = b
        b = fib
    }
    return fib+a
}
