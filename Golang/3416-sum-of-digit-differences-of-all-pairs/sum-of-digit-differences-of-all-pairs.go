func sumDigitDifferences(nums []int) int64 {
    var output int64
    // arrSize := len(nums)
    digitLen := 0
    for temp := nums[0]; temp > 0; temp /= 10 {
        digitLen++
    }
    digitFreq := make([][]int, digitLen)
    for idx := range digitFreq {
        digitFreq[idx] = make([]int, 10)
    }
    // fmt.Println(digitFreq)
    // fmt.Println("Arr Size : ", arrSize)
    for idx, num := range nums {
        n := strconv.Itoa(num)
        for i := 0; i < len(n); i++ {
            digit, _ := strconv.Atoi(string(n[i]))
            output += int64(idx - digitFreq[i][digit])
            digitFreq[i][digit]++
        }
    }
    fmt.Println(digitFreq)
    return output
}
// func findSum(num1, num2 int) int64 {
//     var sum int64
//     for num1 > 0 {
//         digit1 := num1 % 10
//         digit2 := num2 % 10
//         if digit1 != digit2 {
//             sum++
//         }
//         num1 /= 10
//         num2 /= 10
//     }
//     // fmt.Println("num1 : ",num1, "num2 : ", num2, " : ",sum)
//     return sum
// }