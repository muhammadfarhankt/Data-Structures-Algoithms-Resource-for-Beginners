func countTestedDevices(batteryPercentages []int) int {
    count := 0
    for i, n := 0, len(batteryPercentages); i < n; i++ {
        if batteryPercentages[i] - count > 0 {
            count++
        }
    }
    return count
}