func rotate(nums []int, k int)  {
    size := len(nums)
    k = k % size
    for i := 0; i < size / 2; i++ {
        nums[i], nums[size-i-1] = nums[size-i-1], nums[i]
    }
    fmt.Println("Reverse : ", nums)
    for i := 0; i < k/2; i++ {
        nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
    }
    fmt.Println("first : ", nums)
    for i, j := k, 1; i < (size + k) / 2; i++ {
        nums[i], nums[size-j] = nums[size-j], nums[i]
        j++
    }
    fmt.Println("second : ", nums)
}