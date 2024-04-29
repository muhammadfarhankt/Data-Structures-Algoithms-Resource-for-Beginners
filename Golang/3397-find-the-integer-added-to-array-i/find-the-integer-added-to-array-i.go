func addedInteger(nums1 []int, nums2 []int) int {
    sum1, sum2, size := 0, 0, len(nums1)
    for i := 0; i < size; i++ {
        sum1 += nums1[i]
        sum2 += nums2[i]
    }
    return (sum2 - sum1) / size
}