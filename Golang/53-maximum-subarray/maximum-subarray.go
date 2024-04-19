func maxSubArray(nums []int) int {
    maxCurr, maxGlobal := 0, -10001
    for i := 0; i < len(nums); i++ {
        maxCurr += nums[i]
        if maxCurr > maxGlobal {
            maxGlobal = maxCurr
        }
        if maxCurr < 0 {
            maxCurr = 0
        }
    }
    return maxGlobal
}
