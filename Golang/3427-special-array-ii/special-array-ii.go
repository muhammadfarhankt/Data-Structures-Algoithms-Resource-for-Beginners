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
    for idx, query := range queries {
        if parityArr[query[1]] - parityArr[query[0]] > 0 {
            output[idx] = false
        } else {
            output[idx] = true
        }
    }
    return output
}