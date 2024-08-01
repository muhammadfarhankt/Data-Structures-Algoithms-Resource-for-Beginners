func smallestEqual(nums []int) int {
    mod := 101
    for i, num := range nums {
        if i % 10 == num && i < mod {
            // fmt.Println(mod, i, num)
            mod = i
        }
    }
    if mod == 101 {
        return -1
    }
    return mod
}