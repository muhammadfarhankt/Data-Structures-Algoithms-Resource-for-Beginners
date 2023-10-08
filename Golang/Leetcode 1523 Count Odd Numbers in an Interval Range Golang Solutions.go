//Solution 1
func countOdds(low int, high int) int {
    count := 0
    count = (high-low) / 2
    if ((high % 2 != 0) || (low % 2 != 0)  ){
        count++
    }
    return count
}

//Solution 2
func countOdds(low int, high int) int {
    count := 0
    for i:=low; i<=high; i++ {
        if i%2 == 1{
            count++
        }
    }
    return count
}

//Solution 3
func countOdds(low int, high int) int {
    count := 0
    i := 0
    if low%2 == 1 {
        count++
        i = low+2
    } else {
        i = low+1
    }
    for ; i<=high; i+=2 {
            count++
    }
    return count
}
