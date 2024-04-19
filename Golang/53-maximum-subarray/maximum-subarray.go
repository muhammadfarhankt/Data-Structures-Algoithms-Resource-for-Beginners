func maxSubArray(nums []int) int {
	currMax, globalMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
        if currMax < 0 {
			currMax = 0
		}
		currMax += nums[i]
		if globalMax < currMax {
			globalMax = currMax
		}
	}
	return globalMax
}