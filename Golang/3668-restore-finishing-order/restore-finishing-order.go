func recoverOrder(order []int, friends []int) []int {
    friendsArr := [101]int{0}
    output := []int{}
    for _, num := range friends {
        friendsArr[num] = 1
    }
    for _, val := range order {
        if friendsArr[val] == 1 {
            output = append(output, val)
        }
    }
    // fmt.Println(output)
    return output
}