func findTheDifference(s string, t string) byte {
    var sumS, sumT byte
    for i:=0; i<len(s); i++ {
        sumS += s[i]
    }
    for i:=0; i<len(t); i++ {
        sumT += t[i]
    }
    return sumT-sumS
}
