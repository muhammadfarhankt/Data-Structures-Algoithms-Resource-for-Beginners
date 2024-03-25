func findDuplicates(nums []int) []int {
    output := []int{}
    for _, num := range nums {
        idx := abs(num)
        if nums[idx-1] < 0 {
            output = append(output, idx)
        } else {
            nums[idx-1] *= -1
        }
    }
    return output
}
func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func findDuplicates(nums []int) []int {
    size := len(nums)
    // idxArr := make([]int, size)
    output := []int{}
    // for _, num := range nums {
    //     idxArr[num-1]++
    // }
    // for idx, num := range idxArr {
    //     if num == 2 {
    //         output = append(output, idx+1)
    //     }
    // }
    // return output
    numsMap := make(map[int]bool, size)
    for _, num := range nums {
        if numsMap[num] {
            output = append(output, num)
        }
        numsMap[num] = true
    }
    return output
}

func findDuplicates(nums []int) []int {
    size := len(nums)
    idxArr := make([]int, size)
    output := []int{}
    for _, num := range nums {
        idxArr[num-1]++
        if idxArr[num-1] == 2 {
            output = append(output, num)
        }
    }
    // for idx, num := range idxArr {
    //     if num == 2 {
    //         output = append(output, idx+1)
    //     }
    // }
    return output
    // size := len(nums)
    // output := []int{}
    // numsMap := make(map[int]bool, size)
    // for _, num := range nums {
    //     if numsMap[num] {
    //         output = append(output, num)
    //     }
    //     numsMap[num] = true
    // }
    // return output
}
