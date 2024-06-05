func myPow(x float64, n int) float64 {
    if x == 0 {
        return float64(0)
    }
    if n == 0 {
        return float64(1)
    }
    var result float64
    result = math.Pow(x, float64(n))
    return result
}