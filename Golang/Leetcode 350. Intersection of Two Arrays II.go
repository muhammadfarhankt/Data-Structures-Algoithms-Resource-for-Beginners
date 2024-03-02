func intersect(nums1 []int, nums2 []int) []int {
    result := make([]int, 0)
    nums1Map := make(map[int]int, 0)
    for _, num := range nums1 {
        nums1Map[num]++
    }
    for _, num := range nums2 {
        if nums1Map[num] > 0 {
            result = append(result, num)
            nums1Map[num]--
        }
    }
    return result
}
