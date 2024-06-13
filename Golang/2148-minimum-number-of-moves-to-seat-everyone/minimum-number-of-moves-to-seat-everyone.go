func minMovesToSeat(seats []int, students []int) int {
    moves := 0
    sort.Ints(seats)
    sort.Ints(students)
    fmt.Println(seats)
    fmt.Println(students)
    for i := 0; i < len(seats); i++ {
        moves += abs(seats[i] - students[i])
    }
    return moves
}
func abs(num int) int {
    if num >= 0 {
        return num
    }
    return -num
}