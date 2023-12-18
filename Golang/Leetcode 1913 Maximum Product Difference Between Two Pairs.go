func maxProductDifference(nums []int) int {
    sort.Ints(nums)
    n := len(nums)-1
    return ( nums[n] * nums[n-1] - nums[0] * nums[1] )
}
