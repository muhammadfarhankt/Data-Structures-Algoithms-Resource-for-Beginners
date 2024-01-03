func findMatrix(nums []int) [][]int {
    newArr := [][]int{{}}
    row := 1
    sort.Ints(nums)
    newArr[0] = append(newArr[0],nums[0])
    for i := 1; i<len(nums); i++ {
        if nums[i-1] == nums[i] {
            if len(newArr) <= row {
                newArr = append(newArr, []int{nums[i]})
            } else {
                newArr[row] = append(newArr[row],nums[i])
            }
            row++
        } else {
            row = 1
            newArr[0] = append(newArr[0], nums[i])
        }
    }
    return newArr
}
