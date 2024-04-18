func maximumPrimeDifference(nums []int) int {
    left, right := -1, -1
    for i := 0; i < len(nums); i++ {
        if checkPrime(nums[i]) {
            //fmt.Println(nums[i])
            left = i
            break
        }
    }
    for i := len(nums)-1; i >= 0; i-- {
        if checkPrime(nums[i]) {
            //fmt.Println(nums[i])
            right = i
            break
        }
    }
    return right - left
}
func checkPrime(num int) bool {
    if num == 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if num % i == 0 {
            return false
        }
    }
    return true
}