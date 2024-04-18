func maximumPrimeDifference(nums []int) int {
	left, right := -1, -1
	for i, v := range nums {
		if isPrime(v) {
			left = i
			break
		}
	}

	if left == -1 {
		return 0
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if isPrime(nums[i]) {
			right = i
			break
		}
	}

	if right <= left {
		return 0
	}

	return right - left
}

func isPrime(val int) bool {
	if val < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(val))); i++ {
		if val%i == 0 {
			return false
		}
	}

	return true
}