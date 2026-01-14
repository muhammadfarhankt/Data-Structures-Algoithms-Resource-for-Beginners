func absDifference(nums []int, k int) int {
    left_sum, right_sum, length := 0, 0, len(nums)
    if length == 1 {
        return 0
    }
    sort.Ints(nums)
    for i := 0; i < k; i++ {
        left_sum += nums[i]
        right_sum += nums[length - i - 1]
    }
    if (left_sum - right_sum) < 0 {
        return -(left_sum - right_sum)
    }
    return left_sum - right_sum
}