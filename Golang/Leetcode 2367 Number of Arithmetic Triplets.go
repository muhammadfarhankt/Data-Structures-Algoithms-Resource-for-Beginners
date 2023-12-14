func arithmeticTriplets(nums []int, diff int) int {
    count := 0
    numsMap := make(map[int]bool)
    for _, num := range nums {
        numsMap[num] = true
    }
    for _, num := range nums {
        if numsMap[num+diff] {
            if numsMap[num+diff+diff] {
                count++
            }
        }
    }
    return count
}
