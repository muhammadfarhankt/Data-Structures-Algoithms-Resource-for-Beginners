func sortPeople(names []string, heights []int) []string {
    n := len(names)
    heightsMap := make(map[int]string)
    for idx, height := range heights {
        heightsMap[height] = names[idx]
    }
    sort.Ints(heights)
    output := make([]string, 0, n)
    for i := n-1; i >= 0; i-- {
        output = append(output, heightsMap[heights[i]])
    }
    return output
}