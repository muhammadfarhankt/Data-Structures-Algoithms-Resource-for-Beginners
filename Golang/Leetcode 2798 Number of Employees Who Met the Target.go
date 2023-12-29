func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    //sort.Ints(hours)
    //sorting makes greater than O(N) So normal array traversal is enough
    count := 0
    for i:=len(hours)-1; i>=0; i-- {
        if hours[i] >= target {
            count++
        }
    }
    return count   
}
