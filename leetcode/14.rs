enum Solution {}
impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let (car, cdr) = match strs.split_first() {
            Some(words) => words,
            None => {
                return String::new()
            },
        };

        let mut iterables: Vec<_> = cdr
            .iter()
            .map(|s| s.chars())
            .collect();

        car.chars()
           .take_while(|&c| iterables.iter_mut().all(|i| i.next() == Some(c)))
           .collect()
    }
}

fn main() {
    let strs = vec![
        "flower".to_string(), "flow".to_string(), "flight".to_string()
    ];

    println!("{:?}", Solution::longest_common_prefix(strs))
}