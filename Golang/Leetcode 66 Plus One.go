func plusOne(digits []int) []int {
    end := len(digits)-1
    for end >= 0 {
        if digits[end] == 9 {
            digits[end] = 0
            end--
        } else {
            digits[end] += 1
            break
        }
    }
    if digits[0] == 0 {
        digits[0] = 1
        digits = append(digits,0)
    }
    return digits
}
