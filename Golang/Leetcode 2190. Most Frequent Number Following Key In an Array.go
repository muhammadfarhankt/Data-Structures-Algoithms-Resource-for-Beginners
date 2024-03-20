func mostFrequent(nums []int, key int) int {
    numsMap := make(map[int]int)
    maxCount, res := 0, 0
    for i:=0; i < len(nums)-1; i++ {
        if nums[i] == key {
            numsMap[nums[i+1]]++
            if numsMap[nums[i+1]] > maxCount {
                maxCount = numsMap[nums[i+1]]
                res = nums[i+1]
            }
        }
    }
    fmt.Println(numsMap)
    return res
}
