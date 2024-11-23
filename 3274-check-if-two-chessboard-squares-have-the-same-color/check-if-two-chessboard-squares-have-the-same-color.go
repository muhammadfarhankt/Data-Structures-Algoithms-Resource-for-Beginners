    func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
        val1 := (int(coordinate1[0]) - int('a') + int(coordinate1[1]) + 1) 
        val2 := (int(coordinate2[0]) - int('a') + int(coordinate2[1]) + 1)
        if (val1 % 2 == 0 && val2 % 2 == 0) || (val1 % 2 == 1 && val2 % 2 == 1) {
            return true
        }
        return false
    }