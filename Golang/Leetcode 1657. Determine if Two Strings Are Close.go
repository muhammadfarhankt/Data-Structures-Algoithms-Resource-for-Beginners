func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    word1Map := make(map[byte]int)
    word2Map := make(map[byte]int)
    word1Array, word2Array := []int{}, []int{}
    for i, _ := range word1 {
        word1Map[word1[i]]++
    }
    for i, _ := range word2 {
        word2Map[word2[i]]++
        if word1Map[word2[i]] < 1 {
            return false
        }
    }
    for _, value := range word1Map {
        word1Array = append(word1Array, value)
    }
    for _, value := range word2Map {
        word2Array = append(word2Array, value)
    }
    sort.Ints(word1Array)
    sort.Ints(word2Array)
    // if !reflect.DeepEqual(word1Array, word2Array) {
    //     return false
    // }
    for i:=0; i<len(word1Array); i++ {
        if word1Array[i] != word2Array[i] {
            return false
        }
    }
    return true
}
