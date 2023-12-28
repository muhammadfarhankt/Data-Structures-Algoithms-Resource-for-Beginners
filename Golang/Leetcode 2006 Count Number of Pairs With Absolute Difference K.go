func countKDifference(nums []int, k int) int {
    count := 0
    for i:=0; i<len(nums)-1; i++ {
        for j:=i+1; j<len(nums); j++ {
            if (nums[i] - nums[j] == k) || (nums[i] - nums[j] == -k) {
                count++
            }
        }
    }
    return count
}
