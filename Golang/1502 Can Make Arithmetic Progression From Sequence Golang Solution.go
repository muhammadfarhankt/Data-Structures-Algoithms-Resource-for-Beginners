func canMakeArithmeticProgression(arr []int) bool {
    if len(arr) == 2 {
        return true
    }
    sort.Ints(arr)
    difference := arr[1] - arr[0]
    for i:=1; i<len(arr) - 1; i++ {
        if (arr[i+1] - arr[i]) != difference {
            return false
        }
    }
    return true
}
