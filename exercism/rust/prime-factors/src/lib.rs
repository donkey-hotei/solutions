pub fn factors(n: u64) -> Vec<u64> {
    let mut m = n;
    let mut v = vec![];

    for p in (2..).take_while(|p| p * p <= n) {
        while m % p == 0 {
            v.push(p);

            m /= p;
        }
    }

    if 1 < m { v.push(m) }

    v
}
