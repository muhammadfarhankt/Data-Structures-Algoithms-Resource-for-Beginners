func maximumOddBinaryNumber(s string) string {
    // countOne := 0
    // for _, char := range s {
    //     if char == '1' {
    //         countOne++
    //     }
    // }
    // str := ""
    // for i:=1; i<countOne; i++ {
    //     str += "1"
    // }
    // for i:=1; i<=(len(s)-countOne); i++ {
    //     str += "0"
    // }
    // str += "1"
    // return str
    ones, zeroes := make([]byte,0), make([]byte,0)
    for _, char := range s {
        if char == '1' {
            ones = append(ones, '1')
        } else {
            zeroes = append(zeroes, '0')
        }
    }
    if len(ones) == 1 {
        return string(zeroes) + "1"
    }
    ones = ones[:len(ones)-1]
    return string(ones) + string(zeroes) + "1"
}
