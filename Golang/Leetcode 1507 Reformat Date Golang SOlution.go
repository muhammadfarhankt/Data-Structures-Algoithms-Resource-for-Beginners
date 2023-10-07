//Leetcode 1507 Reformat Date Golang SOlution
func reformatDate(date string) string {
    monthMap := map[string]string{"Jan" : "01","Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06", "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12"}
    dateData := strings.Split(date," ")
    day, month, year := dateData[0][:len(dateData[0])-2], dateData[1], dateData[2]
    if len(day) == 1 {
        day = "0"+day
    }
    return year + "-" + monthMap[month] + "-" + day
}
