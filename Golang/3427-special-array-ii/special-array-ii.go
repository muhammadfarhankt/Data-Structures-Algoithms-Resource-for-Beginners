func isArraySpecial(nums []int, queries [][]int) []bool {
    size := len(nums)
    parityArr := make([]int, size)
    for i := 1; i < size; i++ {
        parityArr[i] = parityArr[i-1]
        if nums[i] % 2 == nums[i-1] % 2 {
            parityArr[i]++
        }
    }
    output := make([]bool, len(queries))
    for i := 0; i < len(queries); i++ {
        if parityArr[queries[i][1]] - parityArr[queries[i][0]] == 0 {
            output[i] = true
        }
    }
    return output
}