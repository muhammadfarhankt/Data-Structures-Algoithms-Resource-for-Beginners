func smallestNumber(n int, t int) int {
	for {
		prod := 1
		for temp := n; temp > 0; temp /= 10 {
			prod *= temp % 10
		}
		if prod % t == 0 {
			return n
		} else {
			n++
		}
	}
	return 0
}