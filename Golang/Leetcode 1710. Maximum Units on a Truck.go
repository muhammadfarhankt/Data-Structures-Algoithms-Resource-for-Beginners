func maximumUnits(boxTypes [][]int, truckSize int) int {
    output := 0
    sort.Slice(boxTypes, func(i, j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    for i:=0; i<len(boxTypes); i++ {
        if truckSize >= boxTypes[i][0] {
            output += boxTypes[i][0] * boxTypes[i][1]
            truckSize -= boxTypes[i][0]
        } else {
            output += truckSize * boxTypes[i][1]
            return output
        }
    }
    return output
}
