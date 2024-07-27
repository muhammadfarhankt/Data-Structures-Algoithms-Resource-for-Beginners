func rotate(nums []int, k int)  {
    size := len(nums)
    newArr := make([]int, size)
    k = k % size
    for i := 0; i < k; i++ {
        newArr[i] = nums[size-k+i]
    }
    for i := k; i < size; i++ {
        newArr[i] = nums[i-k]
    }
    fmt.Println(newArr)
    for i := 0; i < size; i++ {
        nums[i] = newArr[i]
    }
}