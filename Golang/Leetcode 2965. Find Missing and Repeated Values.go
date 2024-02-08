func findMissingAndRepeatedValues(grid [][]int) []int {
    l := len(grid) * len(grid)
    numsArr := make([]bool, l+1)
    repeat, sum := 0, 0
    for _, row := range grid {
        for _, num := range row {
            sum += num
            if numsArr[num] {
                repeat = num
            }
            numsArr[num] = true
        }
    }
    missing := (l * (l+1) / 2) - sum + repeat
    return []int{repeat, missing}
}

//Using Hashmap

func findMissingAndRepeatedValues(grid [][]int) []int {
    numsMap := make(map[int]bool)
    repeat, sum := 0, 0
    for _, val := range grid {
        for _, num := range val {
            sum += num
            if numsMap[num] {
                repeat = num
            }
            numsMap[num] = true
        }
    }
    l := len(grid) * len(grid)
    missing := (l * (l+1) / 2) - sum + repeat
    return []int{repeat, missing}
}
