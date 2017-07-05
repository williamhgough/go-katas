package romanNumerals

import "strings"

// Numeral holds a decimal value and string representation of runes.
type Numeral struct {
	Value uint16
	Rune  string
}

var numerals = []Numeral{
	{Value: 10, Rune: "X"},
	{Value: 9, Rune: "IX"},
	{Value: 5, Rune: "V"},
	{Value: 4, Rune: "IV"},
	{Value: 1, Rune: "I"},
}

// IntToNumerals will return roman numeral value of an integer.
func IntToNumerals(num uint16) (roman string) {
	for _, numeral := range numerals {
		amount, remaining := DivMod(num, numeral.Value)
		num = remaining
		roman += strings.Repeat(numeral.Rune, int(amount))
	}
	return
}

// DivMod divides a number and returns its remainder.
func DivMod(number uint16, divisor uint16) (quotient, remainder uint16) {
	quotient = number / divisor
	remainder = number % divisor
	return
}

// NumeralsToInts takes a roman numeral and returns an integer.
func NumeralsToInts(input string) (value uint16) {
	return 0
}
