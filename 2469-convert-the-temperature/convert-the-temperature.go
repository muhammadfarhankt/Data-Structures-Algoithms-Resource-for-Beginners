func convertTemperature(celsius float64) []float64 {
    kelvin := celsius + 273.15
    fahrenheit := celsius * 1.80 + 32.00
    arr := []float64{kelvin,fahrenheit}
    return arr
}