func getFinalState(nums []int, k int, multiplier int) []int {
    for k != 0 {
        idx := minim(nums)
        nums[idx] *= multiplier
        k--
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