func isBalanced(num string) bool {
    var odd byte
    var even byte
    for i := 0; i < len(num); i++ {
        if i % 2 == 0 {
            even += num[i] - '0'
        } else {
            odd += num[i] - '0'
        }
    }
    // fmt.Println("odd : ",odd,"\teven : ",even)
    if odd == even {
        return true
    }
    return false
}