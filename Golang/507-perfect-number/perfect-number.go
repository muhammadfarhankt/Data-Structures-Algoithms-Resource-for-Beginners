func checkPerfectNumber(num int) bool {
    sum, i := 0, 1
    for i <= num/2 {
        if num % i == 0 {
            sum += i
        }
        i++
    }
    if sum == num {
        return true
    }
    return false
}