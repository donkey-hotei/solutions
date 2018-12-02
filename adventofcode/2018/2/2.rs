use std::fs;
use std::collections::HashMap;

fn letter_counts(word: &str) -> HashMap<u8, isize> {
    let mut counts = HashMap::with_capacity(26);

    for letter in word.bytes() {
        *counts.entry(letter).or_insert(0) += 1;
    }

    counts
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let box_ids: Vec<&str> = input.split("\n").collect();

    let mut twice  = 0;
    let mut thrice = 0;

    for box_id in box_ids.iter() {
        let counts = letter_counts(box_id);

        let twice_times  = counts.iter().any(|(_, &count)| count == 2);
        let thrice_times = counts.iter().any(|(_, &count)| count == 3);

        if twice_times {
            twice += 1;
        }

        if thrice_times {
            thrice += 1;
        }
    }

    println!("{:?}", thrice * twice)
}
