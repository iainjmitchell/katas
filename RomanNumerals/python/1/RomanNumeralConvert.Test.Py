import unittest

class RomanNumerals:
    def convert(self, arabic):
        return "I"


class RomanNumeralConvertorTest(unittest.TestCase):
    """Test roman numeral conversion"""

    def test_1_returns_I(self):
        """Test that arabic 1 returns numeral 'I'"""
        roman_numerals = RomanNumerals()
        self.assertEqual(roman_numerals.convert(1), "I")


if __name__ == "__main__":
    unittest.main() 