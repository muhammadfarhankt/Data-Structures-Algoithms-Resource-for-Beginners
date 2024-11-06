func getSneakyNumbers(nums []int) []int {
    numsArr := make([]int, 101, 101)
    output := make([]int, 0, 2)
    for _, num := range nums {
        numsArr[num]++
        if numsArr[num] == 2 {
            output = append(output, num)
        }
    }
    return output
}