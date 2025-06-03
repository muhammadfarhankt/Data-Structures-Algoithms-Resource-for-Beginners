func returnToBoundaryCount(nums []int) int {
    pos, count := 0, 0
    for _, num := range nums {
        pos += num
        if pos == 0 {
            count++
        }
    }
    return count
}