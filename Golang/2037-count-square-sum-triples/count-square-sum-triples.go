func countTriples(n int) int {
    count := 0
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            k := int(math.Sqrt(float64(i*i + j*j)))
            if k <= n && k * k == i * i + j * j {
                count++
            } 
        }
    }
    return count
}