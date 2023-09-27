func createTargetArray(nums []int, index []int) []int {
    targetArray := []int{}
    for i,_ := range index{
        num := nums[i]
        idx := index[i]
        if idx < len(targetArray){
            targetArray = append(targetArray[0:idx],append([]int{num},targetArray[idx:]...)...)    
        } else {
            targetArray = append(targetArray[0:i],num)
        }
    }
    return targetArray
}
