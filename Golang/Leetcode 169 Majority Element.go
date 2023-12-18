func majorityElement(nums []int) int {
    numMap := make(map[int]int)
    for _, num := range nums {
        numMap[num]++
    }
    large := 1
    num := nums[0]
    for index, value := range numMap {
        if value > large {
            large = value
            num = index
        }
    }
    return num
}
