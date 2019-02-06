/*
 * leetcode #64: Minimum Path Sum
 */
use std::cmp;

fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
    let (m, n) = (grid.len(), grid[0].len());

    let mut grid = grid;
    for i in 1..m {
        grid[i][0] += grid[i - 1][0];
    }

    for j in 1..n {
        grid[0][j] += grid[0][j - 1];
    }

    for i in 1..m {
        for j in 1..n {
            grid[i][j] += cmp::min(grid[i - 1][j], grid[i][j - 1]);
        }
    }

    grid[m - 1][n - 1]
}


fn main() {
    let m = vec![
        vec![1, 3, 1],
        vec![1, 5, 1],
        vec![4, 2, 1]
    ];

    assert_eq!(min_path_sum(m), 7);

    println!("tests pass.");
}