func isPowerOfThree(n int) bool {
    // if n < 1 {
    //     return false
    // }
    // for n % 3 == 0 {
    //     n /= 3
    // }
    // return n == 1
    if n == 1 {
        return true
    }
    if n % 3 != 0 || n <= 0 {
        return false
    }
    return isPowerOfThree(n/3)
}
