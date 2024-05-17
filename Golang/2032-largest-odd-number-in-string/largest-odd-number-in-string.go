func largestOddNumber(num string) string {
    size := len(num)
    for i := size - 1; i >= 0; i-- {
        if int(num[i]) % 2 == 1 {
            fmt.Println(num[:i])
            return num[:i+1]
        }
    }
    return ""
}