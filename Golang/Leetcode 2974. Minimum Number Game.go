func numberGame(nums []int) []int {
    slices.Sort(nums)
    size := len(nums)
    // result := make([]int, size)
    for i:=0; i<size; i=i+2 {
        nums[i], nums[i+1] = nums[i+1], nums[i]
    }
    return nums
}
