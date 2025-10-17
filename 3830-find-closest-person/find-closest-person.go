func findClosest(x int, y int, z int) int {
    if x == y {
        return 0
    }
    if abs(z - x) == abs(z - y) {
        return 0
    } else if abs(z - x) > abs(z - y) {
        return 2
    } else {
        return 1
    }
    return 0
}
func abs (x int) int {
    if x < 0 {
        return -x
    }
    return x
}