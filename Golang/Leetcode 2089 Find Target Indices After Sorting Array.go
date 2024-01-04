func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    output := []int{}
    for idx, val := range nums {
        if val == target {
            output = append(output, idx)
        }
    }
    return output
}
