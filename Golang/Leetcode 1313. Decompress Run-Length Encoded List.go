func decompressRLElist(nums []int) []int {
    arr := []int{}
    for i:=0; i<len(nums); i=i+2 {
        for j:=1; j<=nums[i]; j++ {
            arr = append(arr,nums[i+1])
        }
    }
    return arr
}
