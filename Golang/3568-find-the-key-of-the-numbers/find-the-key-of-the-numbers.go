func generateKey(num1 int, num2 int, num3 int) int {
    output := [4]int{}
    result := 0
    for i := 1; i <= 4; i++ {
        output[4-i] = minim(num1 % 10, num2 % 10, num3 % 10)
        num1 /= 10
        num2 /= 10
        num3 /= 10
    }
    for _, digit := range output {
        result = result * 10 + digit
    }
    return result
}

func minim(num1, num2, num3 int) int {
    if num1 < num2 {
        if num1 < num3 {
            return num1
        } else {
            return num3
        }
    } else if num2 < num3 {
        return num2
    }
    return num3
}