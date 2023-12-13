func minPartitions(n string) int {
    large := n[0]
    for i:=1; i<len(n); i++ {
        if large < n[i] {
            large = n[i]
        }
    }
    return int(large - '0')
}
