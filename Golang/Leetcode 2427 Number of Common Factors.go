func commonFactors(a int, b int) int {
    factors := 1
    secondLarge := 0
    if a > b {
        secondLarge = b
    } else {
        secondLarge = a
    }
    for i:=2; i <= secondLarge; i++ {
        if a % i == 0 && b % i == 0 {
            factors++
        }
    }
    return factors
}
