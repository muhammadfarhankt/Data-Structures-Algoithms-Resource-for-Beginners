func groupThePeople(groupSizes []int) [][]int {
    numMap := make(map[int][]int)
    result := [][]int{}
    for i, num := range groupSizes {
        numMap[num] = append(numMap[num],i)
        if len(numMap[num]) == num {
            result = append(result,numMap[num])
            numMap[num] = []int{}
        }
    }
    return result
}
