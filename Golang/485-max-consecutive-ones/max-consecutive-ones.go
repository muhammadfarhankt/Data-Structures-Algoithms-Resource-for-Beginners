func findMaxConsecutiveOnes(nums []int) int {
    count, maxCount := 0, 0
    for _, num := range nums {
        if num == 1 {
            count++
        } else {
            count = 0
        }
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}