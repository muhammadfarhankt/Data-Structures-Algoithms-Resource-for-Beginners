func judgeCircle(moves string) bool {
    x, y := 0, 0
    for _, move := range moves {
        if move == 'U' {
            y++
        } else if move == 'D' {
            y--
        } else if move == 'R' {
            x++
        } else if move == 'L' {
            x--
        }
    }
    if x == 0 && y == 0 {
        return true
    }
    return false
}
