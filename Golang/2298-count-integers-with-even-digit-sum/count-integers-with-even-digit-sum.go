func countEven(num int) int {
    count := 0
    for i := 2; i <= num; i++ {
        sum, temp := 0, i
        for temp > 0 {
            sum += temp % 10
            temp /= 10
        }
        if sum % 2 == 0 {
            count++
        }
    }
    return count
}