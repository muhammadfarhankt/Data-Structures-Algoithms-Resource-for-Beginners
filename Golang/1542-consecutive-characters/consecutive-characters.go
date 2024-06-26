func maxPower(s string) int {
    count, maxCount := 1, 1
    for i := 0; i < len(s)-1; i++ {
        if s[i] == s[i+1] {
            count++
        } else {
            if count > maxCount {
                maxCount = count
            }
            count = 1
        }
    }
    if count > maxCount {
        maxCount = count
    }
    return maxCount
}