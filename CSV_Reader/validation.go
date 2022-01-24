package CSV_Reader

// pH is a decimal value between 0 - 14
func Check_pH(value float64) (pH float64) {
	if value >= 0 && value <= 14 {
		return value
	}
	return -1
}

// Temperature is a celsius value between -50 and 100
func Check_Temperature(value float64) (temperature float64) {
	if value >= -50 && value <= 100 {
		return value
	}
	return -1
}

// Rainfall is a positive number between 0 and 500
func Check_Rainfall(value float64) (rainfall float64) {
	if value >= 0 && value <= 500 {
		return value
	}
	return -1
}
