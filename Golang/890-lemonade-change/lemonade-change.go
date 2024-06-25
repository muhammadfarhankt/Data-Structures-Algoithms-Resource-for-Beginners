func lemonadeChange(bills []int) bool {
    five, ten, twenty := 0, 0, 0
    for _, bill := range bills {
        if bill == 5 {
            five++
        } else if bill == 10 {
            if five > 0 {
                five--
                ten++
            } else {
                return false
            }
        } else {
            if five > 0 && ten > 0 {
                five--
                ten--
                twenty++
            } else if five > 2 {
                five -= 3
            } else {
                return false
            }
        }
    }
    return true
}