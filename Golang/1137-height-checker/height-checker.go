func heightChecker(heights []int) int {
    temp := make([]int, 0, len(heights)) 
    temp = append(temp, heights...)
    count := 0
    sort.Ints(temp)
    fmt.Println(temp)
    fmt.Println(heights)
    for i := 0; i < len(heights); i++ {
        if heights[i] != temp[i] {
            count++
        }
    }
    return count
}