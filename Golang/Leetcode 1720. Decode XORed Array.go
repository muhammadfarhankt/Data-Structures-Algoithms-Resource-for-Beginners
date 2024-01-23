func decode(encoded []int, first int) []int {
    arr := make([]int, len(encoded)+1)
    arr[0] = first
    for i:=0; i<len(encoded); i++ {
        arr[i+1] = encoded[i] ^ arr[i]
    }
    return arr
}
