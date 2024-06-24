func minimumAverage(nums []int) float64 {
    n := len(nums)
    var minAvg float64
    sort.Ints(nums)
    //fmt.Println(nums)
    //fmt.Println(avgArr)
    minAvg = float64(nums[0] + nums[n-1]) / float64(2)
    for i := 1; i < (n/2); i++ {
        avg := float64(nums[i] + nums[n-1-i]) / float64(2)
        if avg < minAvg {
          minAvg = avg
        }
    }
    return minAvg
    
}