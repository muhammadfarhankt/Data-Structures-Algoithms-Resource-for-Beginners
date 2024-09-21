func merge(nums1 []int, m int, nums2 []int, n int)  {
    // if m == 0 {
    //     nums1 = nums2
    // } else if n == 0 {
    //     return
    // }
    l := m + n - 1
    for n > 0 {
        if m > 0 && (nums1[m-1] >= nums2[n-1]) {
            nums1[l] = nums1[m-1]
            m--
            l--
        } else {
            nums1[l] = nums2[n-1]
            n--
            l--
        }
    }
}