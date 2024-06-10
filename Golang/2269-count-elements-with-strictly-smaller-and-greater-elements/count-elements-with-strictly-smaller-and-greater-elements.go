func countElements(nums []int) int {
    sort.Ints(nums)
    min, max := nums[0], nums[len(nums)-1]
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > min && nums[i] < max{
            count++
        }
    }  
    return count
}