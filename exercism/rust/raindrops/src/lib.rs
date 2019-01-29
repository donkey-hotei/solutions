pub fn raindrops(n: u32) -> String {
    if n % 3 != 0 &&
       n % 5 != 0 &&
       n % 7 != 0 { n.to_string() }
    else {
        (1..=n)
            .filter(|x| n % x == 0)
            .fold(String::new(), |mut s, factor| {
                match factor {
                    3 => { s.push_str("Pling"); s },
                    5 => { s.push_str("Plang"); s },
                    7 => { s.push_str("Plong"); s },
                    _ => s
                }
            })
    }
}
