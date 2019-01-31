use std::collections::HashSet;
use rayon::iter::{IntoParallelIterator, ParallelIterator};

pub fn find(n: u32) -> HashSet<[u32; 3]> {
    (0..n / 3)
        .into_par_iter()
        .map(|a| {
            let b = ((n - a) * (n - a) - a * a) / (2 * (n - a));
            let c = n - a - b;

            if a < b && b < c         &&
               a * a + b * b == c * c &&
               a + b + c == n { Some([a, b, c]) }
            else {
                None
            }
        })
        .filter_map(|x| x)
        .collect()
}
