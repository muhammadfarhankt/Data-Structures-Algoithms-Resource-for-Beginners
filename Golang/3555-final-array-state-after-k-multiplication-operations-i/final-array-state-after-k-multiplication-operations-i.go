func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 1; i <= k; i++ {
        idx := minim(nums)
        nums[idx] *= multiplier
    }
    return nums
}
func minim(nums []int) int {
    low := 0
    for i, num := range nums {
        if num < nums[low] {
            low = i
        }
    }
    return low
}