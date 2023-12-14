func minimumOperations(nums []int) int {
    count := 0
    numMap := make(map[int]bool)
    for _, num := range nums {
        if numMap[num] == false && num != 0 {
            count++
        }
        numMap[num] = true
    }
    return count
}
