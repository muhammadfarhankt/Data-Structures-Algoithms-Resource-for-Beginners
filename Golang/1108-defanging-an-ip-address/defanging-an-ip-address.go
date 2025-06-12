func defangIPaddr(address string) string {
    output := ""
    for _, char := range address {
        if char == '.' {
            output = output + "[.]"
        } else {
            output = output + string(char)
        }   
    }
    return output
}