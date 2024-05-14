func divideArray(nums []int) bool {
    numsArr := make([]int, 501, 501)
    for _, num := range nums {
        numsArr[num]++
    }
    for i := 1; i < 501; i++ {
        if numsArr[i] % 2 != 0 {
            return false
        }
    }
    return true
}