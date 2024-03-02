func getCommon(nums1 []int, nums2 []int) int {
    for i,j := 0, 0; i<len(nums1) && j<len(nums2); {
        if nums1[i] == nums2[j] {
            return nums1[i]
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    } 
    return -1
    // nums1Map := make(map[int]bool, 0)
    // for _, num := range nums1 {
    //     nums1Map[num] = true
    // }
    // for _, num := range nums2 {
    //     if nums1Map[num] {
    //         return num
    //     }
    // }
    // return -1
}
