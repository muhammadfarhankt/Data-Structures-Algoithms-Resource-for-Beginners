func maximumValue(strs []string) int {
    maxVal := 0
    for _, str := range strs {
        val, err := strconv.Atoi(str)
        if err != nil {
            val = len(str)
        }
        if val > maxVal {
            maxVal = val
        }
    }
    return maxVal
}
