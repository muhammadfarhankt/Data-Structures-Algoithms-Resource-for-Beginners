func maximumPrimeDifference(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        leftPrime := checkPrime(nums[left])
        rightPrime := checkPrime(nums[right])
        if leftPrime && rightPrime {
            return right - left
        }
        if leftPrime == false {
            left++
        }
        if rightPrime == false {
            right--
        }
    }
    return 0
}
func checkPrime(num int) bool {
    if num == 1 {
        return false
    }else if num == 2 || num == 3 {
        return true
    }
    for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if num % i == 0 {
            return false
        }
    }
    return true
}