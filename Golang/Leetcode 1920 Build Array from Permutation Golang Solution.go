func buildArray(nums []int) []int {
    newArray := make( []int, len(nums))
    for i:=0; i<len(nums); i++ {
        newArray[i] = nums[nums[i]]
    }
    return newArray
}


// second solution

func buildArray(nums []int) []int {
    newArray := []int{}
    for _, val := range nums {
        newArray = append(newArray,nums[val])
    }
    return newArray
}
