package layout

import "strconv"

func fromStringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func fromStringToFloat(s string) float64 {
	value, _ := strconv.ParseFloat(s, 64)
	return value
}

func fromFloat64ToString(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}
