func findLucky(arr []int) int {
    lucky := -1
    numsMap := make(map[int]int, len(arr))
    for _, num := range arr {
        numsMap[num]++
    }
    for idx, val := range numsMap {
        if idx == val {
            if idx > lucky {
                lucky = idx
            }
        }
    }
    return lucky
}