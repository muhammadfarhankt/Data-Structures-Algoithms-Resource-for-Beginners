func findLatestTime(s string) string {
    str := []byte(s)
    if str[0] == '?' {
        if str[1] == '0' || str[1] == '?' || str[1] == '1' {
            str[0] = '1'
        } else {
            str[0] = '0'
        }
    }
    if str[1] == '?' {
        if str[0] == '0' {
            str[1] = '9'
        } else if str[0] == '1' {
            str[1] = '1'
        }
    }
    if str[3] == '?' {
        str[3] = '5'
    }
    if str[4] == '?' {
        str[4] = '9'
    }
    return string(str)
}