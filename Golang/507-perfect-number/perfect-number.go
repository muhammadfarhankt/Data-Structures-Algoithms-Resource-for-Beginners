func checkPerfectNumber(num int) bool {
    if num == 1 {
        return false
    }
    sum := 0
    for i:=1; i*i <= num; i++ {
        if num % i == 0 {
            sum += i + (num / i)
        }
    }
    //fmt.Println(sum)
    if sum-num == num {
        return true
    }
    return false
}