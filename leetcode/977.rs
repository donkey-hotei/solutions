/* leetcode #977: squares of a sorted array */

/*
 * Computes the Vec<i32> that contains of the squares of `a` in non-decreasing order.
 */
fn sorted_squares(a: Vec<i32>) -> Vec<i32> {
    let mut b: Vec<i32> = a
        .iter()
        .map(|x| x * x)
        .collect(); // O(n)

    b.sort(); // O(n log n)
    b
}

fn main() {
    let v = vec![-4, -1, 8, 3, 10];
    // let v = vec![1, 2, 3, 4, 5];
    dbg!(sorted_squares(dbg!(v)));
}