func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    for i, num := range nums {
        if bits.OnesCount(uint(i)) == k {
            sum += num
        }
    }
    return sum
}
// func sumIndicesWithKSetBits(nums []int, k int) int {
//     l := len(nums)
//     sum := 0
//     for i:=0; i<l; i++ {
//         if countBinary(i) == k {
//             sum += nums[i]
//         }
//     }
//     return sum
// }
// func countBinary(num int) int {
//     count := 0
//     for num > 0 {
//         if num % 2 == 1 {
//             count++
//         }
//         num /= 2
//     }
//     return count
// }
