func unequalTriplets(nums []int) int {
    tripletCount := 0
    for i:=0; i<len(nums)-2; i++ {
        for j:=i+1; j<len(nums)-1; j++ {
            if nums[i] == nums[j] { continue }
            for k:=j+1; k<len(nums); k++ {
                if nums[i] != nums[j] && nums[j] != nums[k] && nums[i] != nums[k] {
                    tripletCount++
                }
            }
        }
    }
    return tripletCount
}
