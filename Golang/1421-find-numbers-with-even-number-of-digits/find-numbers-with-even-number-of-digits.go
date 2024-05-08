func findNumbers(nums []int) int {
    evenDigits := 0
    for _, num := range nums {
        if num > 9 && num < 100 || num > 999 && num < 10000 || num == 100000 {
            evenDigits++
        }
    }
    return evenDigits
}