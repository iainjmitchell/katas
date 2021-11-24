fn to_roman(_arabic: &i32, _mappings: &Vec<ArabicToNumeralConversion>) -> String {
    let mut _remainder = _arabic.to_owned();
    let numerals: Vec<String> = _mappings.iter()
        .map(|mapping| {
            let occurrences = (_remainder / mapping.arabic) as usize;
            if occurrences > 0 {
                _remainder = _remainder % mapping.arabic;
            }
            return mapping.numeral.repeat(occurrences);
        })
        .collect();
    return numerals.join("");
}

struct ArabicToNumeralConversion {
    arabic: i32,
    numeral: String
} 

fn get_mapping() -> Vec<ArabicToNumeralConversion> {
    return vec!{
        ArabicToNumeralConversion{arabic: 1000, numeral: String::from("M")},
        ArabicToNumeralConversion{arabic: 900, numeral: String::from("CM")},
        ArabicToNumeralConversion{arabic: 500, numeral: String::from("D")},
        ArabicToNumeralConversion{arabic: 400, numeral: String::from("CD")},
        ArabicToNumeralConversion{arabic: 100, numeral: String::from("C")},
        ArabicToNumeralConversion{arabic: 90, numeral: String::from("XC")},
        ArabicToNumeralConversion{arabic: 50, numeral: String::from("L")},
        ArabicToNumeralConversion{arabic: 40, numeral: String::from("XL")},
        ArabicToNumeralConversion{arabic: 10, numeral: String::from("X")},
        ArabicToNumeralConversion{arabic: 9, numeral: String::from("IX")},
        ArabicToNumeralConversion{arabic: 5, numeral: String::from("V")},
        ArabicToNumeralConversion{arabic: 4, numeral: String::from("IV")},
        ArabicToNumeralConversion{arabic: 1, numeral: String::from("I")}
    };
}

#[cfg(test)]
extern crate test_case;

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test_case(1, String::from("I") ; "1 returns I")]
    #[test_case(4, String::from("IV") ; "4 returns IV")]
    #[test_case(5, String::from("V") ; "5 returns V")]
    #[test_case(9, String::from("IX") ; "9 returns IX")]
    #[test_case(10, String::from("X") ; "10 returns X")]
    #[test_case(40, String::from("XL") ; "40 returns XL")]
    #[test_case(50, String::from("L") ; "50 returns L")]
    #[test_case(90, String::from("XC") ; "90 returns XC")]
    #[test_case(100, String::from("C") ; "100 returns C")]
    #[test_case(400, String::from("CD") ; "400 returns CD")]
    #[test_case(500, String::from("D") ; "500 returns D")]
    #[test_case(900, String::from("CM") ; "900 returns CM")]
    #[test_case(1000, String::from("M") ; "1000 returns M")]
    fn single_numeral_tests(arabic: i32, numeral: String){
        assert_eq!(to_roman(&arabic, &get_mapping()), numeral);
    }

    #[test_case(2, String::from("II"); "2 returns II")]
    #[test_case(3, String::from("III"); "3 returns III")]
    #[test_case(20, String::from("XX"); "20 returns XX")]
    fn two_returns_ii(arabic: i32, numeral: String){
        assert_eq!(to_roman(&arabic, &get_mapping()), numeral);
    }

    #[test_case(1066, String::from("MLXVI"); "1066 returns MLXVI")]
    #[test_case(1977, String::from("MCMLXXVII"); "1977 returns MCMLXXVII")]
    fn well_known_dates(arabic: i32, numeral: String){
        assert_eq!(to_roman(&arabic, &get_mapping()), numeral);
    }
}

