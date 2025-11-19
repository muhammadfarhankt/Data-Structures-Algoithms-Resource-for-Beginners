func missingNumber(nums []int) int {
    arraySum := 0
    l := len(nums)
    for _, num := range nums {
        arraySum += num
    }
    return ((l * (l+1)) / 2) - arraySum
}