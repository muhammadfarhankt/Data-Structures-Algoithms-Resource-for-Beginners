func removeZeros(n int64) int64 {
    var newNum int64
    var count int64
    count = 1
    for n > 0 {
        if n % 10 != 0 {
            newNum = newNum + (count * (n % 10))
            count = count * 10
        }
        n = n / 10
    }
    return newNum
}