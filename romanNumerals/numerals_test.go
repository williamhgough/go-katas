package romanNumerals

import "testing"
import "github.com/stretchr/testify/assert"

func TestNumerals(t *testing.T) {
	assert.Equal(t, "I", Numerals(1), "Returns roman numeral for 1 = I")
	assert.Equal(t, "II", Numerals(2), "Returns roman numeral for 2 = II.")
	assert.Equal(t, "not found", Numerals(53), "Returns not found for values outside of map.")
}
