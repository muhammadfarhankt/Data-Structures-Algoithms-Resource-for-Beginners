func numberOfBeams(bank []string) int {
    laserCount, prev := 0, 0
    rLength, cLength := len(bank), len((bank[0]))
    for row := 0; row < rLength; row++ {
        count := 0
        for col := 0; col < cLength; col++ {
            if bank[row][col] == '1' {
                count++
            }
        }
        if count > 0 {
            laserCount += prev * count
            prev = count
        }
    }
    return laserCount
}
