//array solution (optimised because we know character length is 26
func canConstruct(ransomNote string, magazine string) bool {
    ransomArr := make([]int,26)
    magazineArr := make([]int,26)
    for _, val := range ransomNote {
        ransomArr[(int(val - 'a'))]++
    }
    for _, val := range magazine {
        magazineArr[(int(val - 'a'))]++
    }
    fmt.Println(ransomArr)
    fmt.Println(magazineArr)
    for i, val := range ransomArr {
        if val > magazineArr[i] {
            return false
        }
    }
    return true
}

//hashmap solution
func canConstruct(ransomNote string, magazine string) bool {
    ransomMap := make(map[rune]int,0)
    magazineMap := make(map[rune]int,0)
    for _, val := range ransomNote {
        ransomMap[val]++
    }
    for _, val := range magazine {
        magazineMap[val]++
    } 
    for key, val := range ransomMap {
        if magazineMap[key] < val {
            return false
        }
    }
    return true
}
