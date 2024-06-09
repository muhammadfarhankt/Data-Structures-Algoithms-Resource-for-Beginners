func numberOfChild(n int, k int) int {
    k = k % (2 * (n - 1))
    if k < n {
        return k
    }
    return (2 * (n - 1)) - k
}