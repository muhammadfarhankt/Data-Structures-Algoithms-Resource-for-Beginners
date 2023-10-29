func numIdenticalPairs(nums []int) int {
    count := 0
    numsMap := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        numsMap[nums[i]]++
    }
    for _, element := range numsMap {
		//fmt.Println("i : ",i,"element ; ",element)
		count += element * (element-1) / 2
	}
    return count
}
