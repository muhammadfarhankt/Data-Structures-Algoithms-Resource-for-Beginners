func finalPositionOfSnake(n int, commands []string) int {
    pos := 0
    for _, command := range commands {
        if command == "RIGHT" {
            pos++
        } else if command == "LEFT" {
            pos--
        } else if command == "UP" {
            pos -= n
        } else {
            pos += n
        }
    }
    return pos
}