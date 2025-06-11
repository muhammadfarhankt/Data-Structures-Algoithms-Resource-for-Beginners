func findGCD(nums []int) int {
    small, large := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < small {
            small = nums[i]
        }
        if nums[i] > large {
            large = nums[i]
        }
    }
    for i := small; i >= 1; i-- {
        if large % i == 0 && small % i == 0 {
            return i
        }
    }
    return 1
}