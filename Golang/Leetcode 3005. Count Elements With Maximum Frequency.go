func maxFrequencyElements(nums []int) int {
    numsMap := make(map[int]int)
    count, max := 0, 0
    for _, num := range nums {
        numsMap[num]++
        if numsMap[num] > max {
            max = numsMap[num]
            count = max
        } else if numsMap[num] == max {
            count += max
        }
    }
    return count
}
