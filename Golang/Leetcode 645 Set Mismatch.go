func findErrorNums(nums []int) []int {
    newNums := []int{}
    sum := 0
    n := len(nums)
    sort.Ints(nums)
    for i:=0; i<len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            newNums = append(newNums, nums[i])
        }
        sum += nums[i]
    }
    sum += nums[n-1]
    // for i:=len(nums)-1; i>0; i-- {
    //     if (i+1) != nums[i] {
    //         newNums = append(newNums, (i+1))
    //         break
    //     }
    // }
    newNums = append(newNums, ((n*(n+1)/2)-sum+newNums[0]) )
    return newNums
}
