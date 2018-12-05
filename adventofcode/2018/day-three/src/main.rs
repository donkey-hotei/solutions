use std::fs;
use regex::Regex;

fn main() {
    let input = fs::read_to_string("input.txt").expect("error reading file");
    let regex = Regex::new(
        r"(#\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)")
        .unwrap();

    let mut matrix = [[0u8; 1000]; 1000];

    for line in input.lines() {
        for cap in regex.captures_iter(&line) {
            let left = &cap
                .name("left")
                .unwrap()
                .as_str()
                .parse::<usize>()
                .unwrap();
            let top = &cap
                .name("top")
                .unwrap()
                .as_str()
                .parse::<usize>()
                .unwrap();
            let width = &cap
                .name("width")
                .unwrap()
                .as_str()
                .parse::<usize>()
                .unwrap();
            let height = &cap
                .name("height")
                .unwrap()
                .as_str()
                .parse::<usize>()
                .unwrap();

            for i in *left..left + width {
                for j in *top..top + height {
                    matrix[i][j] += 1;
                }
            }
        }
    }

    let mut overclaims = 0;
    for i in 0..1000 {
        for j in 0..1000 {
            if matrix[i][j] > 1 {
                overclaims += 1;
            }
        }
    }

    println!("{:?}", overclaims)
}
