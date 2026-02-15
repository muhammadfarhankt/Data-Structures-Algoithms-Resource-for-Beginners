func toggleLightBulbs(bulbs []int) []int {
    lightArr := make([]int, 100)
    for _, num := range bulbs {
        if lightArr[num-1] == 0 {
            lightArr[num-1] = 1
        } else {
            lightArr[num-1] = 0
        }
    }
    output := make([]int, 0, 100)
    for i := 0; i < 100; i++ {
        if lightArr[i] == 1 {
            output = append(output, i+1)
        }
    }
    return output
}