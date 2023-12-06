func differenceOfSum(nums []int) int {
    elementSum := 0
    digitSum := 0
    for _,value := range nums {
        elementSum += value
        for value>=1 {
            digitSum = digitSum + value%10
            value /= 10
        }
    }
    return((elementSum - digitSum))
}
