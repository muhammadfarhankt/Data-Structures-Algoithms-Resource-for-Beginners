func alternatingSum(nums []int) int {
    sum, digit := 0, 1
    for i := 0; i < len(nums); i++ {
        sum += digit * nums[i]
        digit *= -1
    }
    return sum
}