func minSteps(s string, t string) int {
    diff := 0
    stMap := make(map[byte]int)
    //sMap := make(map[byte]int)
    //tMap := make(map[byte]int)
    for i, _ := range s {
        stMap[s[i]]++
    }
    for i, _ := range t {
        stMap[t[i]]--
    }
    for _, value := range stMap {
        if value > 0 {
            diff += value
        } else {
            diff -= value
        }
    }
    // for key, value := range sMap {
    //     if value > tMap[key] {
    //         diff += value - tMap[key]
    //     }
    // }
    // for key, value := range tMap {
    //     if value > sMap[key] {
    //         diff += value - sMap[key]
    //     }
    // }
    return diff
}
