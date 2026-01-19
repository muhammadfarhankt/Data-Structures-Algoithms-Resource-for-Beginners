func countSeniors(details []string) int {
    count := 0
    for _, detail := range details {
        age, _ := strconv.Atoi(detail[11:13])
        if age > 60 {
            count++
        }
    }
    return count
}