package romanNumerals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumerals(t *testing.T) {
	assert.Equal(t, "I", Numerals(1), "Returns roman numeral for 1 = I")
	assert.Equal(t, "II", Numerals(2), "Returns roman numeral for 2 = II.")
	assert.Equal(t, "not found", Numerals(53), "Returns not found for values outside of map.")
}

func TestNumeralsToInts(t *testing.T) {
	assert.Equal(t, 1, NumeralsToInts("I"), "Returns integer for roman numeral I = 1")
	assert.Equal(t, 2, NumeralsToInts("II"), "Returns integer for roman numeral II = 2")
	assert.Equal(t, 4, NumeralsToInts("IV"), "Returns integer for roman numeral IV = 4")
	assert.Equal(t, 5, NumeralsToInts("V"), "Returns integer for roman numeral V = 5")
	assert.Equal(t, 6, NumeralsToInts("VI"), "Returns integer for roman numeral VI = 6")
	assert.Equal(t, 7, NumeralsToInts("VII"), "Returns integer for roman numeral VII = 7")
	assert.Equal(t, 9, NumeralsToInts("IX"), "Returns integer for roman numeral IX = 9")
	assert.Equal(t, 9, NumeralsToInts("VIIII"), "Returns integer for roman numeral VIIII = 9")
	assert.Equal(t, 11, NumeralsToInts("XI"), "Returns integer for roman numeral XI = 11")
}
