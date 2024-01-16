type RandomizedSet struct {
    // Map to store value-index pairs for fast lookups and removals
    numsArrIdxMap map[int]int
    // Array to hold the actual elements for random access
    numsArray []int
}

func Constructor() RandomizedSet {
    // Initialize the data structures
    return RandomizedSet{
        numsArrIdxMap: make(map[int]int),
        numsArray:     make([]int, 0),
    }
}

func (this *RandomizedSet) Insert(val int) bool {
    // Check if the value already exists
    if _, ok := this.numsArrIdxMap[val]; ok {
        return false
    }

    // Add the value to the array and map its index
    this.numsArrIdxMap[val] = len(this.numsArray)
    this.numsArray = append(this.numsArray, val)
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    // Check if the value exists
    if idx, ok := this.numsArrIdxMap[val]; ok {
        // Swap the element with the last one to simplify removal
        n := len(this.numsArray)
        this.numsArray[idx] = this.numsArray[n-1]
        this.numsArrIdxMap[this.numsArray[n-1]] = idx  // Update the swapped element's index

        // Remove the last element and update maps
        this.numsArray = this.numsArray[:n-1]
        delete(this.numsArrIdxMap, val)
        return true
    }
    return false
}

func (this *RandomizedSet) GetRandom() int {
    // Generate a random index and return the element at that index
    idx := rand.Intn(len(this.numsArray))
    return this.numsArray[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
