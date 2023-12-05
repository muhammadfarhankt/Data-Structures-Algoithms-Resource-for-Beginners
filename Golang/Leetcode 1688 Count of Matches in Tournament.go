func numberOfMatches(n int) int {
    matches := 0
    for n > 1 {
        matches += n / 2
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = (n - 1) / 2 + 1
        }
    } 
    return matches
}

//func numberOfMatches(n int) int {
  // return n-1
// }
