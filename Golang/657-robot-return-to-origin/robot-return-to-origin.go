func judgeCircle(moves string) bool {
    X, Y := 0, 0
    for _, char := range moves {
        if char == 'U' {
            Y++
        } else if char == 'D' {
            Y--
        } else if char == 'R' {
            X++
        } else {
            X--
        }
    }
    // fmt.Println("x = ",X,"\ny = ",Y)
    if X == 0 && Y == 0 {
        return true
    }
    return false
}