use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let freq: i32 = input
        .split("\n")
        .map(|freq| freq.parse::<i32>())
        .filter_map(Result::ok)
        .sum();

    println!("{:?}", freq)
}