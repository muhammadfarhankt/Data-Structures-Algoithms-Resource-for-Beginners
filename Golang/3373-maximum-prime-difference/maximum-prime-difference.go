func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func maximumPrimeDifference(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		var isLeftPrime = isPrime(nums[left])
		var isRightPrime = isPrime(nums[right])
		if isLeftPrime && isRightPrime {
			return right - left
		}
		if !isLeftPrime {
			left++
		}
		if !isRightPrime {
			right--
		}
	}
	return 0
}