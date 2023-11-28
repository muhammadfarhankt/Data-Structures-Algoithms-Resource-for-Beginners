func smallerNumbersThanCurrent(nums []int) []int {
    returnArray := []int{}
    	for i:=0; i<len(nums); i++ {
	    count := 0
	    for j:=0; j<len(nums); j++ {
	        if nums[i] > nums[j] {
	            count++
	        }
	    }
	    returnArray = append(returnArray,count)
	}
    return returnArray
}
