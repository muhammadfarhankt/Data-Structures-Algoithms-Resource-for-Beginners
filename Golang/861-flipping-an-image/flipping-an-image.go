func flipAndInvertImage(image [][]int) [][]int {
    length := len(image)
    for i := 0; i < length; i++ {
        for j := 0; j < length/2; j++ {
            image[i][j], image[i][length-1-j] = image[i][length-1-j], image[i][j]
        }
    }
    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
            if image[i][j] == 0 {
                image[i][j] = 1
            } else {
                image[i][j] = 0
            }
        }
    }
    return image
}