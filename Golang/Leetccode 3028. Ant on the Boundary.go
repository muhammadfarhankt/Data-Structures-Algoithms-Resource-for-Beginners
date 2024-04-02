func returnToBoundaryCount(nums []int) int {
    currPos, count := 0, 0
    for _, val := range nums {
        currPos += val
        if currPos == 0 {
            count++
        }
    }
    return count
}
