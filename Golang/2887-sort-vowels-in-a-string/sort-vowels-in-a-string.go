func sortVowels(s string) string {
    lowerArr, higherArr := make([]int, 26), make([]int, 26)
    lowerCount, higherCount := 0, 0
    lowerMap := map[rune]bool{ 'a': true, 'e': true, 'i': true, 'o': true, 'u': true }
    higherMap := map[rune]bool { 'A': true, 'E': true, 'I': true, 'O': true, 'U': true }
    size := len(s)
    output := make([]rune, 0, size)
    for _, char := range s {
        if lowerMap[char] {
            lowerArr[int(char - 'a')]++
            lowerCount++
        } else if higherMap[char] {
            higherArr[int(char - 'A')]++
            higherCount++
        }
    }
    lower, higher := make([]rune, 0, lowerCount), make([]rune, 0, higherCount)
    for i := 0; i < 26; i++ {
        for lowerArr[i] > 0 {
            lower = append(lower, rune(i + 'a'))
            lowerArr[i]--
        }
        for higherArr[i] > 0 {
            higher = append(higher, rune(i + 'A'))
            higherArr[i]--
        }
    }
    fmt.Println(string(lower))
    fmt.Println(string(higher))
    i, j := 0, 0
    for _, char := range s {
        if higherMap[char] || lowerMap[char] {
            if i < higherCount {
                output = append(output, higher[i])
                i++
            } else {
                output = append(output, lower[j])
                j++
            }
        } else {
            output = append(output, char)
        }
    }
    // fmt.Println(lowerArr)
    // fmt.Println(higherArr)
    // fmt.Println(lowerMap)
    // fmt.Println(higherMap)
    fmt.Println((output))
    return string(output)
}