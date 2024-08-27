    func gcdOfStrings(str1 string, str2 string) string {
        var small string
        var large string
        sSize, lSize := 0, 0
        if len(str1) > len(str2) {
            small = str2
            large = str1
            sSize, lSize = len(str2), len(str1)
        } else {
            small = str1
            large = str2
            sSize, lSize = len(str1), len(str2)
        }
        // fmt.Println("Small : ", small)
        // fmt.Println("size : ", sSize, "\nLarge : ", large, "Size : ", lSize)
        for i := sSize; i > 0; i-- {
            if sSize % i == 0 && lSize % i == 0 {
                // fmt.Println(small[:i])
                temp1 := strings.Repeat(small[:i], sSize / i)
                temp2 := strings.Repeat(small[:i], lSize / i)
                if temp1 == small && temp2 == large {
                    return small[:i]
                }
            }
        }
        // for i := 0; i < len(small); i++ {
        //     fmt.Println("small[:(i+1)] : ", small[0:i+1])
        //     temp := strings.Repeat(small[0:i+1], len(small) / (i+1))
        //     if temp == small {
        //         fmt.Println("small same")
        //         for j := i; j < len(large); j++ {
        //             fmt.Println("large ; ", strings.Repeat(small[0:i+1], j))
        //             fmt.Println("large og : ", large)
        //             if strings.Repeat(small[0:i+1], j) == large {
        //                 fmt.Println("same")
        //                 // fmt.Println("temp : ", temp)
        //                 return small[0:i+1]
        //             }
        //         }

        //     }
        // }
        return ""
    }