// func sortedSquares(nums []int) []int {
//     numsSquare := make([]int,len(nums))
//     for i:=0; i<len(nums); i++ {
//         numsSquare[i] = nums[i] * nums[i]
//     }
//     sort.Ints(numsSquare)
//     return numsSquare
// }

func sortedSquares(nums []int) []int {
    i := len(nums)-1
    left, right := 0, i
    squareArr := make([]int,i+1)
    for left <= right {
        leftSqr := nums[left] * nums[left]
        rightSqr := nums[right] * nums[right]
        if leftSqr > rightSqr {
            squareArr[i] = leftSqr
            left++
        } else {
            squareArr[i] = rightSqr
            right--
        }
        i--
    }
    return squareArr
}
