func averageValue(nums []int) int {
    count, sum := 0, 0
    for _, num := range nums {
        if num % 6 == 0 {
            sum += num
            count++
        }
    }
    if count == 0 {
        return 0
    }
    return sum/count
}
