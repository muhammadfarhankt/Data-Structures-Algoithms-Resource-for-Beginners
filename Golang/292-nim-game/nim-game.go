func canWinNim(n int) bool {
    if n <= 3 {
        return true
    } else if n % 4 == 0 {
        return false
    }
    return true
}