func countCompleteDayPairs(hours []int) int64 {
    seenMap := make(map[int]int)
    count := int64(0)

    for _, hour := range hours {
        rem := hour % 24
        complement := (24 - rem) % 24

        if _, ok := seenMap[complement]; ok {
            count += int64(seenMap[complement])
        }
        seenMap[rem]++
    }

    return count
}
