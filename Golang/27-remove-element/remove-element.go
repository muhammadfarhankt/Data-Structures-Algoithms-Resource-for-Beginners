func removeElement(nums []int, val int) int {
    count := 0
    for _, num := range nums {
        if num != val {
            nums[count] = num
            count++
        }
    }
    return count
}