func reversePrefix(s string, k int) string {
    output := []byte(s)
    for i := 0; i < k/2; i++ {
        output[i], output[k-i-1] = output[k-i-1], output[i]
    }
    return string(output)
}