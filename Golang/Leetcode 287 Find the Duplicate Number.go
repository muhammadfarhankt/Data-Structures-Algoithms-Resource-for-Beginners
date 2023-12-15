func findDuplicate(nums []int) int {
    numMap := make(map[int]bool)
    for _, num := range nums {
        if numMap[num] {
            return num
        }
        numMap[num] = true
    }
    return 0
}
