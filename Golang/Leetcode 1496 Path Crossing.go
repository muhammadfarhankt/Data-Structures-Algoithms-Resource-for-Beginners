func isPathCrossing(path string) bool {
    visitedMap := make(map[[2]int]bool)
    visitedMap[[2]int{0,0}] = true
    x, y := 0, 0
    for _, char := range path {
        if char == 'N' {
            y++
        } else if char == 'S' {
            y--
        } else if char == 'W' {
            x--
        } else if char == 'E' {
            x++
        }
        if visitedMap[[2]int{x,y}] == true {
            return true
        }
        visitedMap[[2]int{x,y}] = true
    }
    return false
}
