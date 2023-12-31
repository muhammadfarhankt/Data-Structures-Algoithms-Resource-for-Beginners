func alternateDigitSum(n int) int {
    sign := 1
    if len(strconv.Itoa(n)) % 2 == 0 {
        sign = -1    
    }
    sum := 0
    for n > 0 {
        sum += sign * (n % 10)
        n /= 10
        sign *= -1
    }
    return sum
}
