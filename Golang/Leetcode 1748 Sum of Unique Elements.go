func sumOfUnique(nums []int) int {
    numsMap := make(map[int]int)
    sum := 0
    for _,val := range nums {
        numsMap[val]++
    }
    for key,val := range numsMap {
        if val == 1 {
            sum += key
        }
    }
    return sum
}
