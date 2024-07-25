func sumOfEncryptedInt(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += encrypt(num)
    }
    return sum
}
func encrypt(num int) int {
    count, large := 0, 0
    for num > 0 {
        digit := num % 10
        if digit > large {
            large = digit
        }
        num /= 10
        count++
    }
    for i:=1; i<count; i++ {
        large = (large * 10) + (large % 10)
    }
    return large
}