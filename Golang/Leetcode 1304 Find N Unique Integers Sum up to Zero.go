func sumZero(n int) []int {
    var resultArray []int
    if(n%2==1){
        resultArray = append(resultArray,0)
    }
    for i:=1;i<=(n/2);i++ {
        resultArray = append(resultArray,i,-i)
    }
    return resultArray
}
