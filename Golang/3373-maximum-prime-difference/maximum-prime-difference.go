func maximumPrimeDifference(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        leftPrime := checkPrime(nums[left])
        rightPrime := checkPrime(nums[right])
        if leftPrime && rightPrime {
            return right - left
        }
        if !leftPrime {
            left++
        }
        if !rightPrime {
            right--
        }
    }
    return 0
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