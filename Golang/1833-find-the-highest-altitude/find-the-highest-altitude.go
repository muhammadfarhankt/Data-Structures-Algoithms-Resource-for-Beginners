func largestAltitude(gain []int) int {
    height, high := 0, 0
    for _, num := range gain {
        height += num
        if height > high {
            high = height
        }
    }
    return high
}