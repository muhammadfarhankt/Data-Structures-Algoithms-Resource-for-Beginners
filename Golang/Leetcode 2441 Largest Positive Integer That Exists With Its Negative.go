func findMaxK(nums []int) int {
    numsMap := make(map[int]int)
    sort.Ints(nums)
    for _, num := range nums {
         numsMap[num]++
    }
    for i:=len(nums)-1; i>0; i-- {
        if _, ok := numsMap[-nums[i]]; ok {
            return nums[i]
        } 
    }
    return -1
}
