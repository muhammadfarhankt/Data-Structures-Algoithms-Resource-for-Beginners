func leftRightDifference(nums []int) []int {
    l := len(nums)
    if l == 1{
        return []int{0}
    }
    sum := make([]int, l)
    leftRightSum := make([][2]int, l)
    leftRightSum[0][0] = 0
    leftRightSum[l-1][1] = 0
    for i:=1; i<l; i++ {
        leftRightSum[i][0] = leftRightSum[i-1][0] + nums[i-1]
    }
    for i:=l-2; i>=0; i-- {
        leftRightSum[i][1] = leftRightSum[i+1][1] + nums[i+1]
    }
    for i:=0; i<l; i++ {
        diff := leftRightSum[i][1] - leftRightSum[i][0]
        if diff < 0 {
            sum[i] = -diff
        } else {
            sum[i] = diff
        }
    }
    return sum
}
