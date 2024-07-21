func minChanges(n int, k int) int {
    if n == k {
        return 0
    }

    count := 0

    for n > 0 || k > 0 {
        bitN := n & 1
        bitK := k & 1

        if bitK == 1 && bitN == 0 {
            return -1
        }

        if bitN == 1 && bitK == 0 {
            count++
        }

        n >>= 1
        k >>= 1
    }

    return count
}