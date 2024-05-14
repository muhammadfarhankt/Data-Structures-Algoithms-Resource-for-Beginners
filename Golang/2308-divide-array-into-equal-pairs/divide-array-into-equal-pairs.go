func divideArray(nums []int) bool {
    numsMap := make(map[int]int, 500)
    for _, num := range nums {
        numsMap[num]++
    }
    for _, val := range numsMap {
        if val %2 != 0 {
            return false
        }
    }
    return true
}