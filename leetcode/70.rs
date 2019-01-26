#![feature(dbg_macro)]
fn naive_climb_stairs(n: usize) -> usize {
    if n < 0 {
        0
    } else if n == 0 {
        1
    } else {
        climb_stairs(n - 1) + climb_stairs(n - 2)
    }
}

fn climb_stairs(n: usize) -> usize {
    let mut memo = vec![0; n + 1];

    if n == 0 {
        return 0
    } else if n == 1 {
        return 1;
    } else if n == 2 {
        return 2;
    }

    memo[1] = 1;
    memo[2] = 2;

    for i in 3..=n {
        memo[i] = memo[i - 1] + memo[i - 2];
    }

    memo[n]
}

fn main() {
    /*
     * There are two ways to sum up to 2:
     *
     *     1 + 1
     * and 2 + 0
     */
    assert_eq!(climb_stairs(2), 2);
    assert_eq!(climb_stairs(3), 3);
    assert_eq!(climb_stairs(4), 5);
}