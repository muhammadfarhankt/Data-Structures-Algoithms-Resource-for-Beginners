func sumOddLengthSubarrays(arr []int) int {
    sum := 0
    l := len(arr)
    for i:=0; i<l; i++ {
        for j:=i; j<l; j+=2 {
            for k := i; k <= j; k++ {
                sum += arr[k]
            }
        }
    }
    return sum
}
