func countEven(num int) int {
    count := 0
    for i:=2; i<=num; i++ {
        if i < 10 && i % 2 == 0 {
            count++
        } else if i > 10 {
            j := i
            digitSum := 0
            for j>0 {
                digitSum += j % 10
                j = j / 10
            }
            if digitSum % 2 == 0 {
                count++
            }
        }
    }
    return count
}
