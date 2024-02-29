func containsNearbyDuplicate(nums []int, k int) bool {
    // for i:=len(nums)-1; i>0; i-- {
    //     for j:=i-1; j>=0; j-- {
    //         if nums[i] == nums[j] && (i-j) <= k {
    //             return true
    //         }
    //     }
    // }
    // return false


    // numsMap := make(map[int][]int)
    // for idx, val := range nums {
    //     numsMap[val] = append(numsMap[val], idx)
    // }
    // for _, arr := range numsMap {
    //     if len(arr) > 1 {
    //         for i:=len(arr)-1; i>0; i-- {
    //             if (arr[i] - arr[i-1]) <= k {
    //                 return true
    //             }
    //         }
    //     }
    // }
    // return false


    numsMap := make(map[int]int)
    for i:=len(nums)-1; i>=0; i-- {
        val, exists := numsMap[nums[i]]
        if exists && (val - i) <= k {
            return true
        }
        numsMap[(nums[i])] = i
    }
    //fmt.Println(numsMap)
    return false
}
