func intersection(nums [][]int) []int {
    numsArr := make([]int, 1001)
    output := make([]int, 0, len(nums[0]))
    for _, arr := range nums {
        for _, num := range arr {
            numsArr[num]++
        }
    }
    for i := 1; i < 1001; i++ {
        if numsArr[i] == len(nums) {
            output = append(output, i)
        }
    }
    return output
}