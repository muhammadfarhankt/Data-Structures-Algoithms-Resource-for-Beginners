func isPowerOfTwo(n int) bool {
    if n == 1 {
        return true
    } else if n % 2 == 1 {
        return false
    } else if n < 0 {
        return false
    }
    for n > 0 {
        if n/2 == 1 && n % 2 == 0 {
            return true
        } else if n % 2 == 1{
            return false
        }
        n = n / 2
    }
    return false
}
