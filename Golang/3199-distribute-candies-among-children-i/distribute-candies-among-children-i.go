func distributeCandies(n int, limit int) int {
    result := 0
    for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			for k := 0; k <= limit; k++ {
				if i+j+k == n {
					result++
				}
			}
		}
	}

	return result
}