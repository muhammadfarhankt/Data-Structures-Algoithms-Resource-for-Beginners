func convertDateToBinary(date string) string {
    // output := ""
    // output += convert(date[:4]) + "-" + convert(date[5:7]) + "-" + convert(date[8:10])
    fmt.Println(convert(date[:4]))
    // fmt.Println((date[5:7]))
    // fmt.Println((date[8:10]))
    // fmt.Println(output)
    return fmt.Sprintf("%v-%v-%v",convert(date[:4]),convert(date[5:7]),convert(date[8:10]))
}

func convert(val string) string {
    res := ""
    num, _ := strconv.Atoi(val)
    for num > 0 {
        res = strconv.Itoa(num % 2) + res
        num /= 2
    }
    return res
}