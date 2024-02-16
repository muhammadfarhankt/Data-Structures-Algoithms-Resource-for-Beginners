func isMonotonic(nums []int) bool {
    isIncreasing, isDecreasing := true, true
    l := len(nums)
    for i:=0; i<l-1; i++ {
        if nums[i] > nums[i+1] {
            isIncreasing = false
        } else if nums[i] < nums[i+1] {
            isDecreasing = false
        }
        if !isIncreasing && !isDecreasing {
            return false
        }
    }
    return isIncreasing || isDecreasing
}
