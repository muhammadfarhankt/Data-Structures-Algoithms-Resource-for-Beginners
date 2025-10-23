func missingMultiple(nums []int, k int) int {
    numsMap := make(map[int]bool, len(nums))
    for _, num := range nums {
        numsMap[num] = true
    }
    for i := 1; i <= 101; i++ {
        if numsMap[k*i] == true {
            continue
        } else {
            return k * i
        }
    }
    return k
}