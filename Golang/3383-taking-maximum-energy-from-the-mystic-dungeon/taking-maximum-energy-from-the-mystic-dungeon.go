func maximumEnergy(energy []int, k int) int {
    maxEnergy := -1001
    size := len(energy)
    for i := k; i < size; i++ {
        val := energy[i] + energy[i-k]
        if val > energy[i] {
            energy[i] = val
        }
    }
    for i := size - 1; i >= size - k; i-- {
        if energy[i] > maxEnergy {
            maxEnergy = energy[i]
        }
    }
    return maxEnergy
}