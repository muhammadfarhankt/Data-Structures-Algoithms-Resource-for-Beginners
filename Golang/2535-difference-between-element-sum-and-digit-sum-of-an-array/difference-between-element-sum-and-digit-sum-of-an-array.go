func differenceOfSum(nums []int) int {
    digit_sum, element_sum := 0, 0   
    for _, num := range nums {
        element_sum += num
        for num > 0 {
            digit_sum += num % 10
            num /= 10
        }
    }
    return abs_diff(digit_sum, element_sum)
}
func abs_diff(num1, num2 int) int {
    if num1 > num2 {
        return num1 - num2
    }
    return num2 - num1
}