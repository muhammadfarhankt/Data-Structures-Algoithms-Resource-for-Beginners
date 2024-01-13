func minSteps(s string, t string) int {
    diff := 0
    sMap := make(map[byte]int)
    tMap := make(map[byte]int)
    for i, _ := range s {
        sMap[s[i]]++
        tMap[t[i]]--
    }
    for key, value := range sMap {
        if value > tMap[key] {
            diff += value - tMap[key]
        }
    }
    return diff
}
