func rearrangeArray(nums []int) []int {
    size := len(nums)
    result := make([]int, size, size)
    for i, pos, neg := 0, 0, 1; i < size; i++ {
        if nums[i] < 0 {
            result[neg] = nums[i]
            neg += 2
        } else {
            result[pos] = nums[i]
            pos += 2
        }
    }
    return result
}