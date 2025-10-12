func sumDivisibleByK(nums []int, k int) int {
    numsMap := make(map[int]int, len(nums))
    sum := 0
    for _, num := range nums {
        numsMap[num]++
    }
    for idx, val := range numsMap {
        if val % k == 0 {
            sum += idx * val
        }
    }
    return sum
}