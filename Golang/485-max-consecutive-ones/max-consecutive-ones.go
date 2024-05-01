func findMaxConsecutiveOnes(nums []int) int {
    count, maxCount := 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            count += 1
        } else {
            if count > maxCount {
                maxCount = count
            }
            count = 0
        }
    }
    if count > maxCount {
        maxCount = count
    }
    return maxCount
}