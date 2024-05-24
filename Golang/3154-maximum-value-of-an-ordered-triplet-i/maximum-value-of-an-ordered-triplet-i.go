func maximumTripletValue(nums []int) int64 {
    max := int64((nums[0] - nums[1]) * nums[2])
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            for k := j+1; k < len(nums); k++ {
                temp := int64((nums[i] - nums[j]) * nums[k])
                if temp > max {
                    max = temp
                }
            }
        }
    }
    if max < 0 {
        return 0
    }
    return max
}