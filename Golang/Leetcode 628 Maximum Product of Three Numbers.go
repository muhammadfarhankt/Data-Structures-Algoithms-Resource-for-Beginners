func maximumProduct(nums []int) int {
    sort.Ints(nums)
    length := len(nums)-1
    return max(nums[length] * nums[length-1] * nums[length-2], nums[0] * nums[1] * nums[length])
}
