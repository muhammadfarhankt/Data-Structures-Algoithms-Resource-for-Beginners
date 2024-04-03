func areaOfMaxDiagonal(dimensions [][]int) int {
    maxArea, longDiagonal  := 0, 0
    for i:=0; i<len(dimensions); i++ {
        len, bre := dimensions[i][0] , dimensions[i][1]
        diagonal := len * len + bre * bre
        //fmt.Println("longDiagonal ; ",longDiagonal,"\ndiagonal : ",diagonal,"\nArea : ",maxArea,"\n")
        area := len * bre
        if diagonal > longDiagonal {
            longDiagonal = diagonal
            maxArea = len * bre
        } else if diagonal == longDiagonal && area > maxArea {
            maxArea = area
        }
    }
    return maxArea
}
