func findMaxConsecutiveOnes(nums []int) int {
    count, maxCount := 0, 0
    for _, num := range nums {
        if num == 1 {
            count++
            if count > maxCount {
                maxCount = count
            }
        } else {
            count = 0
        }
    }
    return maxCount
}