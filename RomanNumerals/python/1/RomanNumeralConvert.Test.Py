import unittest

class RomanNumerals:
    def __init__(self, arabic_to_numeral_mappings):
        self._arabic_to_numeral_mappings = arabic_to_numeral_mappings

    def convert(self, arabic):
        if arabic in self._arabic_to_numeral_mappings:
            return self._arabic_to_numeral_mappings[arabic]
        else:
            return None


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


if __name__ == "__main__":
    unittest.main() 