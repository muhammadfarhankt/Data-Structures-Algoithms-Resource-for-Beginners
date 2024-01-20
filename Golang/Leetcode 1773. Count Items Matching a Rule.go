func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    count := 0
    ruleKeys := map[string]int{
        "type":0,
        "color":1,
        "name":2,
    }
    rule := ruleKeys[ruleKey]
    for _, item := range items {
        if item[rule] == ruleValue {
            count++
        }
    }
    return count
}
