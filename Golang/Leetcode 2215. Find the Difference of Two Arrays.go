func findDifference(nums1 []int, nums2 []int) [][]int {
    nums1Map := make(map[int]bool)
    nums2Map := make(map[int]bool)
    arr := [][]int{{},{}}
    for _, num := range nums1 {
        nums1Map[num] = true
    }
    for _, num := range nums2 {
        nums2Map[num] = true
    }
    for key, _ := range nums1Map {
        if !nums2Map[key] {
            arr[0] = append(arr[0],key)
        }
    }
    for key, _ := range nums2Map {
        if !nums1Map[key] {
            arr[1] = append(arr[1],key)
        }
    }
    return arr
}
