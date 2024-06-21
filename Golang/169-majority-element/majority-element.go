func majorityElement(nums []int) int {
    n := len(nums)
    numsMap := make(map[int]int, n)
    for i := 0; i < n; i++ {
        numsMap[nums[i]]++
        if numsMap[nums[i]] > (n / 2) {
            return nums[i]
        }
    }
    return 0
}