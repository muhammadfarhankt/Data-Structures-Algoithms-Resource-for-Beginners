func findMaxConsecutiveOnes(nums []int) int {
    count, maxCount := 0, 0
    for i := 0; i < len(nums); i++ {
        count = count * nums[i] + nums[i]
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}