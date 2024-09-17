func uncommonFromSentences(s1 string, s2 string) []string {
    str1, str2 := strings.Split(s1, " "), strings.Split(s2, " ")
    map1, map2 := make(map[string]int, len(str1)), make(map[string]int, len(str2))
    output := make([]string, 0, len(str1))
    for _, str := range str1 {
        map1[str]++
    }
    for _, str := range str2 {
        map2[str]++
    }
    for key, value := range map1 {
        if value == 1 {
            if _, exists := map2[key]; !exists {
                output = append(output, key)
            }
        }
    }
    for key, value := range map2 {
        if value == 1 {
            if _, exists := map1[key]; !exists {
                output = append(output, key)
            }
        }
    }
    return output
}
