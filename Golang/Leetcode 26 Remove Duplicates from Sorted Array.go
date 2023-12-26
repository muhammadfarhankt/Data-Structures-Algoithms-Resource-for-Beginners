func removeDuplicates(nums []int) int {
    insertionPosition := 1
    for i:=1; i<len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[insertionPosition] = nums[i]
            insertionPosition++
        }
    }
    return insertionPosition
}
