func permute(nums []int) [][]int {
    result := make([][]int, 0)
    Permutation(0, nums, &result)
    return result
}
func Permutation(index int, nums []int, ans *[][]int) {
    if index == len(nums) {
        *ans = append(*ans, append([]int{}, nums...))
    }
    for i:=index; i<len(nums); i++ {
        nums[index], nums[i] = nums[i], nums[index]
        Permutation(index+1, nums, ans)
        nums[index], nums[i] = nums[i], nums[index]
    }
}
