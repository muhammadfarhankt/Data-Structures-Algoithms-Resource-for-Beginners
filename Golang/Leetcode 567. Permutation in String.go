func checkInclusion(s1 string, s2 string) bool {
    l1, l2 := len(s1), len(s2)
    if l1 > l2 {
        return false
    }
    s1Array, s2Array := [26]int{}, [26]int{}
    for i:=0; i<l1; i++ {
        s1Array[int(s1[i] - 'a')]++
        s2Array[int(s2[i] - 'a')]++
    }
    // fmt.Println("s1Array", s1Array)
    // fmt.Println("s2Array", s2Array)
    if isEqual(s1Array, s2Array) {
        return true
    }
    for i:=l1; i<l2; i++ {
        s2Array[int(s2[i-l1] - 'a')]--
        s2Array[int(s2[i] - 'a')]++
        // fmt.Println("i : ",i,"s2Array : ",s2Array)
        if isEqual(s1Array, s2Array) {
            return true
        }
    }
    return false
}
func isEqual(arr1, arr2 [26]int) bool {
    for i:=0; i<26; i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}
