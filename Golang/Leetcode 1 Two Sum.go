func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i, num := range nums {
        if _, ok := numsMap[(target - num)]; ok {
            j := numsMap[(target - num)]
            return []int{i,j}
        }
        numsMap[num] = i
    }
    return []int{}
}
