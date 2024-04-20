func countDistinctIntegers(nums []int) int {
    numsMap := make(map[int]bool, len(nums)*2)
    for _, num := range nums {
        numsMap[num] = true
        rev := reverse(num)
        numsMap[rev] = true
        nums = append(nums, rev)
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