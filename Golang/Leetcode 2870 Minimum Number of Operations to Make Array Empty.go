func minOperations(nums []int) int {
    numsMap := make(map[int]int)
    count := 0
    for _, num := range nums {
        numsMap[num]++
    }
    for _, value := range numsMap {
        if value == 1{
                return -1
        } else if value % 3 == 0 {
            count += value / 3
            //val = val % 3
        } else if (value - 2) % 3 == 0 {
            count += (value / 3) + 1
        } else if (value - 4) % 3 == 0 {
            count += (value-4) / 3 + 2
        }else if value % 2 == 0 {
            count += value / 2
        } else {
            return -1
        }
    }
    return count
}
