func losingPlayer(x int, y int) string {
    for turn := 0; ; turn++ {
        if x >= 1 && y >= 4 {
            x -= 1
            y -= 4
        } else {
            if turn % 2 == 0 {
                return "Bob"
            } else {
                return "Alice"
            }
        }
    }
}
