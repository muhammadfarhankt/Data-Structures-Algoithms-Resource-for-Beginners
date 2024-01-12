func halvesAreAlike(s string) bool {
    count := 0
    for i, char := range s {
        if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
            if i < len(s)/2 {
                count++
            } else {
                count--
            }            
        }
    }
    return count == 0
}
