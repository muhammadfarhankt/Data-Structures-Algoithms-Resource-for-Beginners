func dominantIndices(nums []int) int {
    count, length := 0, len(nums)
    avgArr := make([]int, length, length)
    for i, sum := length - 2, 0; i >= 0; i-- {
        sum += nums[i+1]
        avgArr[i] = sum / (length - i - 1)
    }
    for i := 0; i < length - 1; i++ {
        // fmt.Println(nums[i], avgArr[i])
        if nums[i] > avgArr[i] {
            count++
        }
        // fmt.Println(count)
    }
    // fmt.Println(avgArr)
    return count
}