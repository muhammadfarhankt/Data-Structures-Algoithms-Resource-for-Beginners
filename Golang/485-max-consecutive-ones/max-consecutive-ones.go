func findMaxConsecutiveOnes(nums []int) int {
    count, maxCount := 0, 0
    for _, num := range nums {
        count = count * num + num
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}