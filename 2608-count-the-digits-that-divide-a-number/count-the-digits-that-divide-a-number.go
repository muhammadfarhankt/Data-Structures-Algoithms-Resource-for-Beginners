func countDigits(num int) int {
    temp, count := num, 0
    for temp > 0 {
        if num % (temp % 10) == 0 {
            count++
        }
        temp /= 10
    }
    return count
}