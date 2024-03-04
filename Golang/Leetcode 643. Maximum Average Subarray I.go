func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i:=0; i<k; i++ {
        sum += nums[i]
    }
    high := sum
    for i:=k; i<len(nums); i++ {
        sum = sum + nums[i] - nums[i-k]
        if sum > high {
            high = sum
        }
    }
    return float64(high) / float64(k)
}
