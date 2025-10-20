func smallestIndex(nums []int) int {
    for i, num := range nums {
        sum := 0
        for num > 0 {
            sum += num % 10
            num /= 10
        }
        if i == sum {
            return i
        }
    }
    return -1
}