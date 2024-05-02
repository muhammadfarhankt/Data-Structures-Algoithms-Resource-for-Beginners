func rearrangeArray(nums []int) []int {
    size := len(nums)
    positive, negative := make([]int, 0, size/2), make([]int, 0, size/2)
    for _, num := range nums {
        if num < 0 {
            negative = append(negative, num)
        } else {
            positive = append(positive, num)
        }
    }
    fmt.Println(positive)
    fmt.Println(negative)
    for odd, even, i := 0, 0, 0; i < size; i++ {
        if i % 2 == 0 {
            nums[i] = positive[even]
            even++
        } else {
            nums[i] = negative[odd]
            odd++
        }
    }
    return nums
}