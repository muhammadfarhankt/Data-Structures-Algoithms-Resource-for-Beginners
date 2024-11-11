func summaryRanges(nums []int) []string {
    var output []string
    for i := 0; i < len(nums); i++ {
        j := i
        for (j + 1 < len(nums) && nums[j+1] == nums[j] + 1) {
            j++
        }
        if i == j {
            output = append(output, fmt.Sprintf("%v", nums[i]))
        } else {
            output = append(output, fmt.Sprintf("%d->%d", nums[i], nums[j]))
        }
        i = j
    }
    return output
}