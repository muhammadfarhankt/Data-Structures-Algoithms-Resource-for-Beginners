func minElement(nums []int) int {
    temp, sum := nums[0], 0
    for temp > 0 {
        val := temp % 10
        sum += val
        temp /= 10
    }
    min := sum
    for _, num := range nums {
        temp, sum = num, 0
        for temp > 0 {
            val := temp % 10
            sum += val
            temp /= 10
        }
        if min > sum {
            min = sum
        }
    }
    return min
}