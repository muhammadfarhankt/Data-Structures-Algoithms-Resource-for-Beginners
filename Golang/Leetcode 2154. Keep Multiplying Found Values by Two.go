func findFinalValue(nums []int, original int) int {
    numsMap := make(map[int]bool)
    for _, num := range nums {
        if num % original == 0 {
            numsMap[num] = true
        }
    }
    for numsMap[original] {
        original *= 2
    }
    return original
}
