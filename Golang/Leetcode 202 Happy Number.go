func isHappy(n int) bool {
   visitedMap := make(map[int]bool)
   if n == 1 {
       return true
   }
   for visitedMap[n] != true {
       visitedMap[n] = true
       n = sumDigitSquares(n)
       if n == 1 {
           return true
       }
   }
   return false
}
func sumDigitSquares (num int) int {
    sum := 0
    for num > 0{
        sum += (num % 10) * (num % 10)
        num /= 10
    }
    return sum
}
