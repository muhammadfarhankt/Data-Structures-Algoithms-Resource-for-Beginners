func earliestTime(tasks [][]int) int {
    min := tasks[0][0] + tasks[0][1]
    for _, val := range tasks {
        if (val[0] + val[1]) < min {
            min = val[0] + val[1]
        }
    }
    return min
}