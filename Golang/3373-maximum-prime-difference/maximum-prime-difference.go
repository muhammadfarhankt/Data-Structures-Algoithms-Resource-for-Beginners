func maximumPrimeDifference(nums []int) int {
    primesArr := []int{}
    for i, n := range nums {
        if checkPrime(n) {
            primesArr = append(primesArr, i)
        }
    }
    return primesArr[len(primesArr)-1] - primesArr[0]
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