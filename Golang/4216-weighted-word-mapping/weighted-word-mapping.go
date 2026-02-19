func mapWordWeights(words []string, weights []int) string {
    output := ""
    for _, word := range words {
        sum := 0
        for _, char := range word {
            val := int(char - 'a')
            sum += weights[val]
        }
        sum = sum % 26
        mapped := 'z' - rune(sum)
        output += string(mapped)
    }
    return output
}
