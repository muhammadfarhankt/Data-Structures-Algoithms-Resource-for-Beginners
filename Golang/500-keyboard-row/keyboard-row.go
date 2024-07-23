func findWords(words []string) []string {
    first := map[byte]bool{'q': true, 'Q': true, 'w': true, 'W': true, 'e': true, 'E': true, 'r': true, 'R': true, 
    't': true, 'T': true, 'y': true, 'Y': true, 'u': true, 'U': true, 'i': true, 'I': true, 'o': true, 'O': true, 'p': true, 'P': true}
    
    second := map[byte]bool{'a': true, 'A': true, 's': true, 'S': true, 'd': true, 'D': true, 'f': true, 'F': true, 
    'g': true, 'G': true, 'h': true, 'H': true, 'j': true, 'J': true, 'k': true, 'K': true, 'l': true, 'L': true}

    third := map[byte]bool{'z': true, 'Z': true, 'x': true, 'X': true, 'c': true, 'C': true, 'v': true, 'V': true, 'b': true, 'B': true, 'n': true, 'N': true, 'm': true, 'M': true}

    output := make([]string, 0, len(words))
    for _, str := range words {
        flag := false
        for i := 0; i < len(str); i++ {
            if !first[str[i]] {
                flag = true
                break
            }
        }
        if !flag {
            output = append(output, str)
            continue
        }

        flag = false
        for i := 0; i < len(str); i++ {
            if !second[str[i]] {
                flag = true
                break
            }
        }
        if !flag {
            output = append(output, str)
            continue
        }

        flag = false
        for i := 0; i < len(str); i++ {
            if !third[str[i]] {
                flag = true
                break
            }
        }
        if !flag {
            output = append(output, str)
        }
    }

    return output
}