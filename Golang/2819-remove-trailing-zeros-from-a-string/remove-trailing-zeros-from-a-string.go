func removeTrailingZeros(num string) string {
    for i := len(num)-1; i >= 0; i-- {
        if num[i] != '0' {
            return num[:i+1]
        }
    }
    return num
}