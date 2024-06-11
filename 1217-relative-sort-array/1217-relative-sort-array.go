func relativeSortArray(arr1 []int, arr2 []int) []int {
    idxArr1 := make([]int, 1001)
    size1, size2 := len(arr1), len(arr2)
    output := make([]int, 0, (size1 + size2))
    for _, num := range arr1 {
        idxArr1[num]++
    }
    for _, num := range arr2 {
        output = append(output, num)
        for idxArr1[num] > 1 {
            output = append(output, num)
            idxArr1[num]--
        }
        idxArr1[num]--
    }
    for idx, num := range idxArr1 {
        for i := 1; i <= num; i++ {
            output = append(output, idx)
        }
    }
    return output
}