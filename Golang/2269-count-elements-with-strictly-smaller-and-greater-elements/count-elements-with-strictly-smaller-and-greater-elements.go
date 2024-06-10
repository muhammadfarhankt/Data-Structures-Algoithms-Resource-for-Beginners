func countElements(nums []int) int {
    max, min := -100001, 100001
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
        }
        if nums[i] < min {
            min = nums[i]
        }
    }
    fmt.Println("max : ", max, "\tmin : ", min)
    for i := 0; i < len(nums); i++ {
        if nums[i] > min && nums[i] < max{
            count++
        }
    }  
    return count
}