pub fn series(digits: &str, len: usize) -> Vec<String> {
    let mut s = 0;
    let mut e = s + len;
    let mut results = Vec::new();

    while e < digits.len() + 1 {
        results.push(digits[s..e].to_string());

        s += 1;
        e += 1;
    }

    results
}
