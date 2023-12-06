func duplicateZeros(arr []int)  {
    leng := len(arr) - 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			for j := leng; j > i; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
			i = i + 1
		}
	}
}
