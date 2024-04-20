func countDistinctIntegers(nums []int) int {
    numsMap := make(map[int]bool, len(nums)*2)
    for i := 0; i < len(nums); i++ {
        numsMap[nums[i]] = true
        numsMap[reverse(nums[i])] = true
    }
    return len(numsMap)
}
func reverse(num int) int {
    rev := 0
    for num > 0 {
        rev = (rev * 10) + (num % 10)
        num /= 10
    }
    return rev
}