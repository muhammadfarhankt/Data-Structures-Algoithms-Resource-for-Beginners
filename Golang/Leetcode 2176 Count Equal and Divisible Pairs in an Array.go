count := 0
    for i:=0; i<len(nums)-1; i++ {
        for j:=i+1; j<len(nums); j++ {
            if ( nums[i] == nums[j] ) &&  ( (i * j) % k == 0 ) {
                count++
            }
        }
