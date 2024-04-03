func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
   count := 0
   for i:=0; i<len(arr1); i++ {
        flag := true
        for j:=0; j<len(arr2); j++ {
            if abs(arr2[j] - arr1[i]) <= d {
                flag = false
                break
            }
        }
        if flag {
            count++
        }
    }
   return count 
}
func abs(num int) int {
    if num<0 {
        return num*-1
    }
    return num
}
