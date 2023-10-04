func xorOperation(n int, start int) int {
    answer := 0
    for i:=0; i<n; i++ {
        answer ^= start+2*i
    }
    return answer
}
