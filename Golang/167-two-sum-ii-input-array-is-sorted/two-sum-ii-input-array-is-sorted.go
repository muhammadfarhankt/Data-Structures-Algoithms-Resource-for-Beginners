func twoSum(numbers []int, target int) []int {
    output := []int{0, 0}
    size := len(numbers)
    numsMap := make(map[int]int, size)
    for idx, num := range numbers {
        if _, err := numsMap[target - num]; err {
            return []int{numsMap[target - num]+1, idx+1}
        }
        numsMap[num] = idx
    }
    return output
}