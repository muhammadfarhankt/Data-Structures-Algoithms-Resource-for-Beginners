func uniqueOccurrences(arr []int) bool {
    numsMap := make(map[int]int)
    
    // Count occurrences of each value
    for _, num := range arr {
        numsMap[num]++
    }
    
    // Check for unique frequencies
    freqMap := make(map[int]bool)
    for _, val := range numsMap {
        if freqMap[val] {
            return false
        }
        freqMap[val] = true
    }
    return true
}
