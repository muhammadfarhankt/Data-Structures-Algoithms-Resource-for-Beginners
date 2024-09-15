func stableMountains(height []int, threshold int) []int {
    output := make([]int, 0, len(height))
    for i := 1; i < len(height); i++ {
        if height[i-1] > threshold {
            output = append(output, i)
        }
    }
    return output
}