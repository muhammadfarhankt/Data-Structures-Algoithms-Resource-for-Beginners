func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    var median float64
    arr := make([]int, 0, (m+n))
    arr = append(nums1, nums2...)
    sort.Ints(arr)
    fmt.Println(arr)
    if size := len(arr); size % 2 == 1 {
        median = float64(arr[(size/2)])
        fmt.Println(median)
    } else {
        val := float64(arr[size/2]) / 2 + float64(arr[(size/2)-1]) / 2
        fmt.Println("values  :  ", arr[size/2],  arr[(size/2)-1])
        median = float64(val)
        fmt.Println("median : ", median)
    }
    return median
}