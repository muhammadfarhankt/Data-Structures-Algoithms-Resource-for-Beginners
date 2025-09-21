func maxKDistinct(nums []int, k int) []int {
    numsMap := make(map[int]bool, len(nums))
    arr := make([]int, 0, len(nums))
    for _, val := range nums {
        numsMap[val] = true
    }
    for val, _ := range numsMap {
        arr = append(arr, val)
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })
    if k > len(arr) {
        k = len(arr)
    }
    return arr[:k]
}