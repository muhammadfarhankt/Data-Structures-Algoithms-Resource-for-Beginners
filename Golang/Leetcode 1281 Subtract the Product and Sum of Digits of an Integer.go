func subtractProductAndSum(n int) int {
    product, sum := 1, 0
    for n > 0 {
        sum += n % 10
        product *= n % 10
        n /= 10
    }
    return (product - sum)
}
