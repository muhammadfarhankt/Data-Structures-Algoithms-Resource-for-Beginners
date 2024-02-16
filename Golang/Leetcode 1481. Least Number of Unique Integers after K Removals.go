func findLeastNumOfUniqueInts(arr []int, k int) int {
    freqMap := make(map[int]int)
    for _, num := range arr {
        freqMap[num]++
    }
    l := len(freqMap)
    freqArr := make([]int,len(arr)+1)
    for _, freq := range freqMap {
        freqArr[freq]++
    }
    //fmt.Println(freqArr)
    for i:=1; k>0; i++ {
        remove := i * freqArr[i]
        if k >= remove {
            k -= remove
            l -= freqArr[i]
        } else {
            l -= (k /  i)
            return l
        }
    }
    return l
}
