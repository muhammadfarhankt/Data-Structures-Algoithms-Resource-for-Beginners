func mirrorDistance(n int) int {
    if n < 10 {
        return 0
    }
    rev := reverse(n)
    if n - rev < 0 {
        return -1 * (n - rev)
    }
    return n - rev
}
func reverse(n int) int {
    rev := 0
    for n > 0 {
        rev = rev * 10 + n % 10
        n = n / 10
    }
    return rev
}