func averageValue(nums []int) int {
    sum, count := 0, 0
    for _, num := range nums {
        if num % 2 == 0 && num % 3 == 0 {
            sum += num
            count++
        }
    }
    if count == 0 {
        return 0
    }
    return int(sum / count)
}