func interpret(command string) string {
    output := ""
    for i:=0; i<len(command); i++ {
        if command[i] == 'G' {
            output = output + "G"
        } else if command[i] == '(' && command[i+1] == ')' {
            output = output + "o"
            i++
        } else {
            output = output + "al"
            i = i + 3
        }
    }
    return output
}
