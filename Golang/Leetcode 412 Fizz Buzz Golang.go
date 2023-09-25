func fizzBuzz(n int) []string {
    string1 := make([]string,n)
  for i:=1;i<=n;i++ {
        if i%3==0 && i%5==0 {
            string1[i-1] = "FizzBuzz"
        } else if i%3==0 {
            string1[i-1] = "Fizz"
        }  else if i%5==0 {
            string1[i-1] = "Buzz"
        } else {
            string1[i-1] = strconv.Itoa(i)
        }
    }
    return string1
}
