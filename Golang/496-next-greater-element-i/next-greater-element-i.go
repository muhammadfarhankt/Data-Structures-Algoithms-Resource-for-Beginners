func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nums2Map := map[int]int{}
    
    for i, num := range nums2 {
        nums2Map[num] = i
    }
    
    result := make([]int, len(nums1))
    for i, num := range nums1 {
        found := false
        for nextIndex := nums2Map[num] + 1; nextIndex < len(nums2); nextIndex++ {
            if nums2[nextIndex] > num {
                result[i] = nums2[nextIndex]
                found = true
                break
            }
        }
        
        if !found {
            result[i] = -1
        }
    }
    
    return result
}