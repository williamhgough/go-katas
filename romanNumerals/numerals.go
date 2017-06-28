package romanNumerals

// Numerals will return roman numeral value of an integer.
func Numerals(num int) string {
	numerals := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
	}

	for key, val := range numerals {
		if num == key {
			return val
		}
	}
	return "not found"
}
