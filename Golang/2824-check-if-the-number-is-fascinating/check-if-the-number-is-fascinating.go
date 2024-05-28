func isFascinating(n int) bool {
    // num := make([]int, 10)
    // for temp := n; temp > 0; temp /= 10 {
    //     num[(temp%10)]++
    // }
    // for temp := 2*n; temp > 0; temp /= 10 {
    //     num[(temp%10)]++
    // }
    // for temp := 3*n; temp > 0; temp /= 10 {
    //     num[(temp%10)]++
    // }
    // // fmt.Println(num)
    // if num[0] > 0 {
    //     return false
    // }
    // for i := 1; i < 10; i++ {
    //     if num[i] > 1 {
    //         return false
    //     }
    // }
    // return true
    return n == 192 || n == 219 || n == 273 || n == 327
}