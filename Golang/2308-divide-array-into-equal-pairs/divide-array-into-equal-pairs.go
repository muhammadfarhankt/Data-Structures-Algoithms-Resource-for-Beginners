func divideArray(nums []int) bool {
    numsArr := make([]int, 501, 501)
    for _, num := range nums {
        numsArr[num]++
    }
    for _, val := range numsArr {
        if val %2 != 0 {
            return false
        }
    }
    return true
}