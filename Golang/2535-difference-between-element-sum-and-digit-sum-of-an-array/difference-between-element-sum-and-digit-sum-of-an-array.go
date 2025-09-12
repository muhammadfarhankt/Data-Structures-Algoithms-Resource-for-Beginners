func differenceOfSum(nums []int) int {
    elementSum, digitSum := 0, 0
    for _, num := range nums {
        elementSum += num
        for num > 0 {
            digitSum += num % 10
            num /= 10
        }
    }
    if elementSum - digitSum < 0 {
        return -1 * (elementSum - digitSum)
    }
    return (elementSum - digitSum)
}