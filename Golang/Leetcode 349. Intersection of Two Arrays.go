func intersection(nums1 []int, nums2 []int) []int {
    // nums1Map, commonMap := make(map[int]bool, 0), make(map[int]bool, 0)
    // for _, num := range nums1 {
    //     nums1Map[num] = true
    // }
    // for _, num := range nums2 {
    //     if nums1Map[num] {
    //         commonMap[num] = true
    //     }
    // }
    // arr := []int
    // for key, _ := range commonMap {
    //     arr = append(arr, key)
    // }
    // return arr
    result := []int{}
    nums1Map := make(map[int]bool, 0)
    for _, num := range nums1 {
        nums1Map[num] = true
    }
    for _, num := range nums2 {
        if nums1Map[num] {
            result = append(result, num)
            nums1Map[num] = false
        }
    }
    return result
}
