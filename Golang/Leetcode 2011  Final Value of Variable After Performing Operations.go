func finalValueAfterOperations(operations []string) int {
    x := 0
    for _, obj := range operations {
        if obj[1] == 43 {
            x++
        } else {
            x--
        }
    }
    return x
}
