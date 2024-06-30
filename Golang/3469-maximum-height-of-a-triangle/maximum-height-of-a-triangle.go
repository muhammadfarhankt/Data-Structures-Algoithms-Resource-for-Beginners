func maxHeightOfTriangle(red int, blue int) int {
    return max(helper(red, blue), helper(blue, red))
}

func helper(red int, blue int) int {
    h := 0
    i := 1

    for {
        if i % 2 == 1 {
            if red >= i {
                red -= i
            } else {
                break
            }
        } else {
            if blue >= i {
                blue -= i
            } else {
                break
            }
        }
        h++
        i++
    }

    return h
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}