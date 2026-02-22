func subarraySum(nums []int) int {
    sum, n := 0, len(nums)
    for i := 0; i < n; i++ {
        start := max(0, i - nums[i])
        for j := start; j <= i; j++ {
            sum += nums[j]
        }
    }
    return sum
}
func max(num1, num2 int) int {
    if num1 > num2 {
        return num1
    }
    return num2
}