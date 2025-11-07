func findMissingElements(nums []int) []int {
    min, max := nums[0], nums[0]
    output := make([]int, 0, 100)
    numsArr := make([]bool, 101)
    for _, num := range nums {
        if num > max {
            max = num
        } else if num < min {
            min = num
        }
        numsArr[num] = true
    }
    for i := min; i <= max; i++ {
        if !numsArr[i] {
            output = append(output, i)
        }
    }
    return output
}