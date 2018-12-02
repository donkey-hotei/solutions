use std::fs;
use std::collections::HashSet;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let changes = input
        .split("\n")
        .map(|freq| freq.parse::<i32>())
        .filter_map(Result::ok);

    let mut frequency = 0;
    let mut frequencies = HashSet::new();

    for change in changes.cycle() {
        frequency += change;

        if frequencies.contains(&frequency) {
            break
        }

        frequencies.insert(frequency);
    }

    println!("{:?}", frequency)
}