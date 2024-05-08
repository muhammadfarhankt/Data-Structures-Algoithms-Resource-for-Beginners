func findNumbers(nums []int) int {
    evenDigits := 0
    for _, num := range nums {
        if evenLength(num) {
            evenDigits++
        }
    }
    return evenDigits
}
func evenLength(num int) bool {
    length := 0
    for num > 0 {
        num /= 10
        length++
    }
    if length % 2 == 0 {
        return true
    }
    return false
}