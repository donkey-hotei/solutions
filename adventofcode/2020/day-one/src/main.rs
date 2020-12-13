use std::fs;
use std::collections::HashSet;



fn part_one(entries: &[i64]) {
   let mut complements: HashSet<i64> = HashSet::new();

    for entry in entries {
        if complements.contains(entry) {
            println!("{:?}", entry * (2020 - entry));
            break;
        }

        complements.insert(2020 - entry);
    }
}

fn part_two(entries: &[i64]) {
    let mut set: HashSet<i64> = HashSet::new();

    set.extend(entries);

    for a in entries {
        for b in entries {
            if set.contains(&(2020 - a - b)) {
                println!("{:?}", a * b * (2020 - a - b));
                break;
            }
        }
    }
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let entries = input
        .split("\n")
        .map(|entry| entry.parse::<i64>())
        .filter_map(Result::ok)
        .collect::<Vec<_>>();

    part_one(&entries);
    part_two(&entries);
}
