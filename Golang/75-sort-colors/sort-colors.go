func sortColors(nums []int)  {
    zero, one, two := 0, 0, 0
    for _, num := range nums {
        if num == 0 {
            zero++
        } else if num == 1 {
            one++
        } else {
            two++
        }
    }
    i := 0
    for zero > 0 {
        nums[i] = 0
        i++
        zero--
    }
    for one > 0 {
        nums[i] = 1
        i++
        one--
    }
    for two > 0 {
        nums[i] = 2
        i++
        two--
    }
}