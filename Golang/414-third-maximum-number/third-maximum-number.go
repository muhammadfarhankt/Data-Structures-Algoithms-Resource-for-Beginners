func thirdMax(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    } else if n == 2 {
        if nums[0] > nums[1] {
            return nums[0]
        } else {
            return nums[1]
        }
    }
    numsMap := make(map[int]bool, n)
    for _, num := range nums {
        numsMap[num] = true
    }
    mapLen := len(numsMap)
    setArr := make([]int, 0, mapLen)
    for num, _ := range numsMap {
        setArr = append(setArr, num)
    }
    if len(setArr) == 1 {
        return setArr[0]
    } else if len(setArr) == 2 {
        if setArr[0] > setArr[1] {
            return setArr[0]
        } else {
            return setArr[1]
        }
    }
    third, second, large := 0, 0, 0
    if setArr[0] > setArr[1] {
        if setArr[0] > setArr[2] {
            large = setArr[0]
            if setArr[2] > setArr[1] {
                second = setArr[2]
                third = setArr[1]
            } else {
                second = setArr[1]
                third = setArr[2]
            }
        } else {
            large = setArr[2]
            if setArr[0] > setArr[1] {
                second = setArr[0]
                third = setArr[1]
            } else {
                second = setArr[1]
                third = setArr[0]
            }
            second = setArr[0]
            third = setArr[1]
        }
    } else if setArr[1] > setArr[2] {
        large = setArr[1]
        if setArr[2] > setArr[0] {
            second = setArr[2]
            third = setArr[0]
        } else {
            second = setArr[0]
            third = setArr[2]
        }
    } else {
        large = setArr[2]
        if setArr[1] > setArr[0] {
            second = setArr[1]
            third = setArr[0]
        } else {
            second = setArr[0]
            third = setArr[1]
        }
    }
    fmt.Println("Array : ",setArr[0], setArr[1], setArr[2])
    fmt.Println("Large : ", large, "Second : ", second, "third : ", third)
    for i := 3; i < len(setArr); i++ {
        if setArr[i] > large {
            third = second
            second = large
            large = setArr[i]
        } else if setArr[i] > second {
            third = second
            second = setArr[i]
        } else if setArr[i] > third {
            third = setArr[i]
        }
    }
    fmt.Println("Final Large : ", large, "Second : ", second, "third : ", third)
    return third
}