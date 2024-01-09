func thirdMax(nums []int) int {
    sort.Ints(nums)
    if len(nums) < 3 {
        return nums[(len(nums)-1)]
    }
    pos := 1
    for i:=len(nums)-1; i > 0; i-- {
        if nums[i] != nums[i-1] {
            pos++
        }
        if pos == 3 {
            return nums[i-1]
        }
    }
    return nums[(len(nums)-1)]
}
