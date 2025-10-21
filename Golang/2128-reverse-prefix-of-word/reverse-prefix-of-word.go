func reversePrefix(word string, ch byte) string {
    result := ""
    pos := -1
    for i := 0; i < len(word); i++ {
        if word[i] == ch {
            pos = i
            break
        }
    }
    if pos != -1 {
        for i := 0; i <= pos; i++ {
            result += string(word[pos-i])
        }
        fmt.Println(result)
        result += word[pos+1:]
    } else {
        return word
    }
    return result
}