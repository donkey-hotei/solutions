pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    (1..limit)
        .filter(|n| {
            factors
                .iter()
                .filter(|&&p| p != 0)
                .any(|p| n % p == 0)
        })
        .sum()
}
