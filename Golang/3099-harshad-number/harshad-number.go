func sumOfTheDigitsOfHarshadNumber(x int) int {
    sum, temp := 0, x
    for temp > 0 {
        sum += temp % 10
        temp /= 10
    }
    if x % sum == 0 {
        return sum
    }
    return -1
}