func separateDigits(nums []int) []int {
    output := make([]int, 0, len(nums))
    for _, num := range nums {
        count, temp := 0, num
        arr := make([]int, 0, 6)
        for temp > 0 {
            rem := temp % 10
            temp = temp / 10
            arr = append(arr, rem)
            count++
        }
        // fmt.Println(arr)
        for i := 0; i < count/2; i++ {
            temp := arr[i]
            arr[i] = arr[count - i - 1]
            arr[count - i - 1] = temp
        }
        // fmt.Println(arr)
        output = append(output, arr...)
    }
    return output
}