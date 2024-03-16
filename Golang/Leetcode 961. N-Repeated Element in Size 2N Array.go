func repeatedNTimes(nums []int) int {
    numsMap := make(map[int]bool)
    for _, num := range nums {
        if numsMap[num] {
            return num
        }
        numsMap[num] = true
    }
    return 0
}
