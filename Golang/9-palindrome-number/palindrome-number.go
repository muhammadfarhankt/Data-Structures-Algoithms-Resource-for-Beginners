func isPalindrome(x int) bool {
    temp := 0
    newNum := 0
    y := x
    for y>0 {
        temp = y%10
        newNum = newNum*10 + temp
        y=y/10
    }
    if newNum == x {
        return true
    }
    return false
}