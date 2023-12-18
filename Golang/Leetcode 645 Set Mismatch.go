func findErrorNums(nums []int) []int {
    sum, duplicate := 0, 0
    n := len(nums)
    sort.Ints(nums)
    for i:=0; i<n-1; i++ {
        if nums[i] == nums[i+1] {
            duplicate = nums[i]
        }
        sum += nums[i]
    }
    sum += nums[n-1]
    return []int{duplicate, ((n*(n+1)/2)-sum+duplicate) }
}
