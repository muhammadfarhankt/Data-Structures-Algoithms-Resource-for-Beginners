func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		if nBasePalindrome(n, i) == false {
			return false
		}
	}
	return true
}
func nBasePalindrome(num int, base int) bool {
    baseNum := ""
	for num > 0 {
        baseNum += string(num % base)
		num = num / base
	}
    return palindrome(baseNum)
}
func palindrome(num string) bool {
	for i,j :=0, len(num)-1; i<=j; i,j = i+1, j-1 {
        if num[i] != num[j] {
            return false
        }
    }
    return true
}
