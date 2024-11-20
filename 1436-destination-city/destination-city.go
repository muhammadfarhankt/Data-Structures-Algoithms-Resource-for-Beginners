func destCity(paths [][]string) string {
    destMap := make(map[string]string)
    for _, val := range paths {
        destMap[val[0]] = val[1]
    }
    // fmt.Println(destMap)
    // init := paths[0][0]
    // fmt.Println(init)
    for _, val := range destMap {
        if _, exists := destMap[val]; !exists {
            return val
        }
        // init = destMap[init]
        // fmt.Println(init)
    }
    return ""
}