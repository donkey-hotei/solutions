use std::cmp;
use std::collections::HashSet;

enum Solution { }

impl Solution {
    pub fn distribute_candies(candies: Vec<i32>) -> i32 {
        let give = candies.len() / 2;
        let mut seen = HashSet::new();

        for candy in &candies {
            if !seen.contains(&candy) {
                seen.insert(candy);
            }
        }

        cmp::min(seen.len(), give) as i32
    }
}

fn main() {
    let candies = vec![1, 1, 1, 1, 1, 2, 2, 3, 3, 3];
    println!("{:?}", Solution::distribute_candies(candies));
}
