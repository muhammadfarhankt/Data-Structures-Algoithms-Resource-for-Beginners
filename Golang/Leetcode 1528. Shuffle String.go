func restoreString(s string, indices []int) string {
    str := make([]rune,len(s))
    for i, c := range s {
        str[indices[i]] = c
    }
    fmt.Println(str)
    return string(str)
}
