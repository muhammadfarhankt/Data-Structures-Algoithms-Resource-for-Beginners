func sumOfTheDigitsOfHarshadNumber(x int) int {
    digitSum := 0
    temp := x
    for temp > 0 {
        digitSum += temp % 10
        temp /= 10
    }
    if x % digitSum == 0 {
        return digitSum
    }
    return -1
}
