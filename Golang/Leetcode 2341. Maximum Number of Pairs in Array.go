func numberOfPairs(nums []int) []int {
    numsMap := make(map[int]int)
    count, rem := 0, 0
    for _, num := range nums {
        numsMap[num]++
        // if numsMap[num] == 2 {
        //     count++
        //     numsMap[num] = 0
        // }
    }
    // fmt.Println("numsMap : ", numsMap)
    for _, val := range numsMap {
        count += val/2
        rem += val%2
    }
    // fmt.Println("count : ",count,"\nrem : ", rem)
    return []int{count, rem}
}
