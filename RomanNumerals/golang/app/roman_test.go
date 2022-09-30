package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arabicToRomanMappings = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func FuzzStraightConversions(fuzzyTest *testing.F) {
	for testCase := range arabicToRomanMappings {
		fuzzyTest.Add(testCase)
	}
	fuzzyTest.Fuzz(func(test *testing.T, testInput int) {
		var expected string = arabicToRomanMappings[testInput]
		assert.Equal(test, expected, ToRoman(arabicToRomanMappings, testInput))
	})
}

func TestTwoReturnsII(test *testing.T) {
	assert.Equal(test, "II", ToRoman(arabicToRomanMappings, 2))
}

func TestThreeReturnsIII(test *testing.T) {
	assert.Equal(test, "III", ToRoman(arabicToRomanMappings, 3))
}

func TestThirtyReturnsXXX(test *testing.T) {
	assert.Equal(test, "XXX", ToRoman(arabicToRomanMappings, 30))
}

func TestThirtyOneReturnsXXXI(test *testing.T) {
	assert.Equal(test, "XXXI", ToRoman(arabicToRomanMappings, 31))
}

func Test1999ReturnsMCMXCIX(test *testing.T) {
	assert.Equal(test, "MCMXCIX", ToRoman(arabicToRomanMappings, 1999))
}
