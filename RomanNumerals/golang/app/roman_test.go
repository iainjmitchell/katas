package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arabicToRomanMappings = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
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
