func removeAnagrams(words []string) []string {     
    output := []string{}
    for i:=len(words)-1; i>0; i-- {
        if !anagram(words[i], words[i-1]) {
            output = append(output, words[i])
        }
    }
    output = append(output, words[0])
    length := len(output)
    for i:=0; i<length/2; i++ {
        output[i], output[length-i-1] = output[length-i-1], output[i]
    }
    return output
}
func anagram(w1, w2 string) bool {
    if len(w1) != len(w2) {
        return false
    }
    charArray := [26]int{}
    for i, char := range w1 {
        charArray[(int(char - 'a'))]++
        charArray[(int(w2[i] - 'a'))]--
    }
    return charArray == [26]int{}
}
