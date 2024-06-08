func clearDigits(s string) string {
    var result []rune
    
    for _, char := range s {
        if char >= '0' && char <= '9' {
            
            if len(result) > 0 && result[len(result)-1] >= 'a' && result[len(result)-1] <= 'z' {
                result = result[:len(result)-1]
            }
            
        } else {
            result = append(result, char)
        }
    }
    return string(result)
    
}