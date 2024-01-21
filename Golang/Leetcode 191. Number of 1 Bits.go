func hammingWeight(num uint32) int {
    count := 0
    str := fmt.Sprintf("%b", num)
    for _, char := range str {
        if char == '1' {
            count++
        }
    }
    return count
}
