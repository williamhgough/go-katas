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

func divMod(quotient int, remainder int) {

}

// NumeralsToInts takes a roman numeral and returns an integer.
func NumeralsToInts(input string) int {

	var value int

	numerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	switch input {
	case "IV":
		return 4
	case "IX":
		return 9
	}

	for _, char := range input {
		for key, val := range numerals {
			if char == key {
				value += val
			}
		}
	}

	return value
}
