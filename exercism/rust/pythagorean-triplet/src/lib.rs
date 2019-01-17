use std::collections::HashSet;

pub fn find(n: u32) -> HashSet<[u32; 3]> {
    (0..n / 3).flat_map(|a| (a + 1..n / 2).map(move |b| (a, b)))
        .map(|(a, b)| (a, b, n - a - b))
        .filter(|&(a, b, c)| a*a + b*b == c*c)
        .map(|(a, b, c)| [a, b, c])
        .collect()
}
