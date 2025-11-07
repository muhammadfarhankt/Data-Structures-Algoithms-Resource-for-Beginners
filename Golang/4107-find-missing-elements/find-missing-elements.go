func findMissingElements(nums []int) []int {
    min, max := nums[0], nums[0]
    output := make([]int, 0, 100)
    numsMap := make(map[int]bool)
    for _, num := range nums {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
        numsMap[num] = true
    }
    for i := min; i <= max; i++ {
        if !numsMap[i] {
            output = append(output, i)
        }
    }
    return output
}