func maximumStrongPairXor(nums []int) int {
    max := nums[0] ^ nums[0]
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            if mod(nums[i], nums[j]) <= min(nums[i], nums[j]) {
                if (nums[i] ^ nums[j]) > max {
                    max = nums[i] ^ nums[j]
                }
            }
        }
    }
    return max
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func mod(a, b int) int {
    if a - b >= 0 {
        return a - b
    }
    return b - a
}