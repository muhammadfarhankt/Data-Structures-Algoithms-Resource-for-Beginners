func maximumCount(nums []int) int {
    pos, neg := 0, 0
    for _, num := range nums {
        if num < 0 {
            neg++
        } else if num >0 {
            pos++
        }
    }
    if neg > pos {
        return neg
    }
    return pos
}