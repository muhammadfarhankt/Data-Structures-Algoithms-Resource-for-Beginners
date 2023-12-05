func xorOperation(n int, start int) int {
    xor := start
    for i:=1; i < n; i++ {
        xor = xor ^ (start + 2 * i)
    }
    return xor
}
