func groupAnagrams(strs []string) [][]string {
    strMap := make(map[[26]int] []string)
    for _, word := range strs {
        wordArr := strArr(word)
        strMap[wordArr] = append(strMap[wordArr], word)
    }
    output := make([][]string, len(strMap))
    i := 0
    for _, val := range strMap {
        output[i] = val
        i++
    }
    return output
}
func strArr (word string) [26]int  {
    arr := [26]int{}
    for _, char := range word {
        arr[(int(char - 'a'))]++
    }
    return arr
}
