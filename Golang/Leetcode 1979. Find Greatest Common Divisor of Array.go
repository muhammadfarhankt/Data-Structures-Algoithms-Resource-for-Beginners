func findGCD(nums []int) int {
    small, large := 0, 0
    if nums[0] > nums[1] {
        small, large = nums[1], nums[0]
    } else {
        small, large = nums[0], nums[1]
    }
    for i:=0; i<len(nums); i++ {
        if nums[i] < small {
            small = nums[i]
        }
        if nums[i] > large {
            large = nums[i]
        }
    }
    for i:=small; i>0; i-- {
        if small % i == 0 && large % i == 0 {
            return i
        }
    }
    return 1
}
