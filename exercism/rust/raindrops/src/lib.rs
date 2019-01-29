pub fn raindrops(n: u32) -> String {
    if n % 3 != 0 && n % 5 != 0 && n % 7 != 0 {
        return n.to_string()
    }

    (1..=n)
        .filter(|x| n % x == 0)
        .fold(String::new(), |mut string, factor| {
            match factor {
                3 => {
                    string.push_str("Pling");
                    string
                },
                5 => {
                    string.push_str("Plang");
                    string
                },
                7 => {
                    string.push_str("Plong");
                    string
                },
                _ => string
            }
        })
}
