import unittest

class RomanNumerals:
    def __init__(self, arabic_to_numeral_mappings):
        self._arabic_to_numeral_mappings = arabic_to_numeral_mappings

    def convert(self, arabic):
        if arabic in self._arabic_to_numeral_mappings:
            return self._arabic_to_numeral_mappings[arabic]
        else:
            closest_arabic = self._closest_arabic_number_with_numeral(arabic)
            return self.convert(closest_arabic) + self.convert(arabic-closest_arabic)

    def _closest_arabic_number_with_numeral(self, arabic):
        return next(matching_arabic 
            for matching_arabic 
            in self._arabic_to_numeral_mappings 
            if arabic / matching_arabic >= 1)


class RomanNumeralConvertorTest(unittest.TestCase):
    """Test roman numeral conversion"""
    ARABIC_TO_NUMERAL_MAPPINGS = {
        1000: "M",
        900: "CM",
        500: "D",
        100: "C",
        90: "XC",
        50: "L",
        40: "XL",
        10: "X",
        9: "IX",
        5: "V",
        4: "IV",
        1: "I"
    }

    def test_straight_arabic_to_numeral_conversions(self):
        """Test that straight conversions of arabics to numerals work (e.g. 1 returns 'I', 5 returns 'V')"""
        roman_numerals = RomanNumerals(arabic_to_numeral_mappings=self.ARABIC_TO_NUMERAL_MAPPINGS)
        for arabic in self.ARABIC_TO_NUMERAL_MAPPINGS:
            self.assertEqual(roman_numerals.convert(arabic), self.ARABIC_TO_NUMERAL_MAPPINGS[arabic])

    def test_2_returns_II(self):
        """Test that 2 returns 'II'"""
        roman_numerals = RomanNumerals(arabic_to_numeral_mappings=self.ARABIC_TO_NUMERAL_MAPPINGS)
        self.assertEqual(roman_numerals.convert(2), "II")

    def test_3_returns_III(self):
        """Test that 3 returns 'III'"""
        roman_numerals = RomanNumerals(arabic_to_numeral_mappings=self.ARABIC_TO_NUMERAL_MAPPINGS)
        self.assertEqual(roman_numerals.convert(3), "III")

    def test_30_returns_XXX(self):
        """Test that 30 returns 'XXX'"""
        roman_numerals = RomanNumerals(arabic_to_numeral_mappings=self.ARABIC_TO_NUMERAL_MAPPINGS)
        self.assertEqual(roman_numerals.convert(30), "XXX")

    def test_2019_returns_MMXIX(self):
        """Test that 30 returns 'MMXIX'"""
        roman_numerals = RomanNumerals(arabic_to_numeral_mappings=self.ARABIC_TO_NUMERAL_MAPPINGS)
        self.assertEqual(roman_numerals.convert(2019), "MMXIX")

if __name__ == "__main__":
    unittest.main() 