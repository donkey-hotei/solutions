use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let grid = input
        .lines()
        .map(|line| line.as_bytes())
        .collect::<Vec<_>>();

    let n = grid[0].len();

    let dx = [1, 3, 5, 7, 1];
    let dy = [1, 1, 1, 1, 2];

    let mut tree_count: u64 = 1;

    for (dx, dy) in dx.iter().zip(&dy) {
        // begin in the upper left
        let mut x = 0;
        let mut y = 0;

        let mut count = 0;

        while y + dy < grid.len() {
            x = (x + dx) % n;
            y = (y + dy);

            if grid[y][x] == '#' as u8 {
                count += 1;
            }
        }

        tree_count *= count;
    }

    println!("{:?}", tree_count)
}
