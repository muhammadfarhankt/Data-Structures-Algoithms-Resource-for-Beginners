func findAnagrams(s string, p string) []int {
    l1, l2 := len(s), len(p)
    if l1 < l2 {
        return []int{}
    }
    output := []int{}
    sArray, pArray := [26]int{}, [26]int{}
    for i:=0; i<l2; i++ {
        sArray[int(s[i] - 'a')]++
        pArray[int(p[i] - 'a')]++
    }
    if isEqual(sArray, pArray) {
        output = append(output, 0)
    }
    for i:=l2; i<l1; i++ {
        sArray[int(s[i-l2] - 'a')]--
        sArray[int(s[i] - 'a')]++
        if isEqual(sArray, pArray) {
            output = append(output, i-l2+1)
        }
    }
    return output
}
func isEqual(arr1, arr2 [26]int) bool {
    for i:=0; i<26; i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}
