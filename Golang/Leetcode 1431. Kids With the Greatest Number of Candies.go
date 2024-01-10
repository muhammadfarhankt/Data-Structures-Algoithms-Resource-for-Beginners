func kidsWithCandies(candies []int, extraCandies int) []bool {
    large := candies[0]
    for i:=1; i<len(candies); i++ {
        if candies[i] > large {
            large = candies[i]
        }
    }
    boolArray := make([]bool,len(candies),len(candies))
    for i, candy := range candies {
        boolArray[i] = candy+extraCandies >= large
    }
    return boolArray
}
