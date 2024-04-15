func findIntersectionValues(nums1 []int, nums2 []int) []int {
    nums1Arr, nums2Arr := make([]bool, 101, 101), make([]bool, 101, 101)
    answer := make([]int, 2, 2)
    for _, num := range nums1 {
        nums1Arr[num] = true
    }
    for _, num := range nums2 {
        nums2Arr[num] = true
    }
    for _, num := range nums1 {
        if nums2Arr[num] {
            answer[0]++
        }
    }
    for _, num := range nums2 {
        if nums1Arr[num] {
            answer[1]++
        }
    }
    return answer
}