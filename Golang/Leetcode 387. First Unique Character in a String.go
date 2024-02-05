func firstUniqChar(s string) int {
    // charMap := make(map[byte]int, 26)
    // for i:=0; i<len(s); i++ {
    //     charMap[s[i]]++
    // }
    // for i:=0; i<len(s); i++ {
    //     if charMap[s[i]] == 1 {
    //         return i
    //     }
    // }
    charArr := make([]int, 26)
    for _, char := range s {
        charArr[int(char - 'a')]++
    }
    for i, char := range s {
        if charArr[int(char - 'a')] == 1 {
            return i
        }
    }
    return -1
}
