func buildArray(nums []int) []int {
    newArray := make( []int, len(nums))
    for i:=0; i<len(nums); i++ {
        newArray[i] = nums[nums[i]]
    }
    return newArray
}
