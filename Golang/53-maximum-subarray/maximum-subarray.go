func maxSubArray(nums []int) int {
    maxCurr, maxGlobal := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        if maxCurr < 0 {
            maxCurr = 0
        }
        maxCurr += nums[i]
        if maxCurr > maxGlobal {
            maxGlobal = maxCurr
        }
    }
    return maxGlobal
}
