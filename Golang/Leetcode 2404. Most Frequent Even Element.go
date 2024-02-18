func mostFrequentEven(nums []int) int {
    numsMap := make(map[int]int)
    for _, num := range nums {
        if num % 2 == 0 {
            numsMap[num]++
        }
    }
    evenElement := -1
    count := 0
    for key, val := range numsMap {
        if val > count || (val == count && key < evenElement) {
            count = val
            evenElement = key
        }
    }
    return evenElement
}
