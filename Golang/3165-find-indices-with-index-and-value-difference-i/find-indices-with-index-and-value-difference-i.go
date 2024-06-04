func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i + indexDifference; j < n; j++ {
            if abs(nums[i] - nums[j]) >= valueDifference {
                return []int{i, j}
            }
        }
    }
    return []int{-1, -1}
}
func abs(num int) int {
    if num >= 0 {
        return num
    }
    return -num
}