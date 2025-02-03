func smallestNumber(n int) int {
    x := n
    for {
        if x % 2 == 0 {
            x++
            continue
        }
        temp := x
        flag := false
        for temp > 0 {
            if temp % 2 == 0 {
                flag = true
                break
            }
            temp /= 2
        }
        if flag == false {
            return x
        }
        x++
    }
    return n
}