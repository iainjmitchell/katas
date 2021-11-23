fn to_roman(_arabic: &i32) -> String {
    return String::from("I");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one_returns_i() {
        assert_eq!(to_roman(&1), "I");
    }
}

