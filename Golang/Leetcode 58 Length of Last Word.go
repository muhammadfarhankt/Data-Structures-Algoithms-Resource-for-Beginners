func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    result := strings.Split(s," ")
    return len(result[len(result)-1])
}
