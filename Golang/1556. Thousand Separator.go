func thousandSeparator(n int) string {
    numString := strconv.Itoa(n)
    if n < 1000 {
        return numString
    }
    for i:=len(numString)-3; i>0; i-=3 {
        numString = numString[:i]+"."+numString[i:]
    }
    return numString
} 
