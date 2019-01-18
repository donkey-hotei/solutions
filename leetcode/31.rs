fn next_permutation(a: &mut Vec<i32>) {
    let n = a.len();

    if let Some(i) = (1..n).rev().find(|&i| a[i - 1] < a[i]) {
        let j = (1..n).rev().find(|&j| a[i - 1] < a[j])
            .unwrap();

        a.swap(i - 1, j);
        a[i..].reverse();
    } else {
        a.reverse();
    }
}

fn main() {
    let mut a = vec![1, 2, 3];
    next_permutation(&mut a);
    println!("{:?}", a);
    next_permutation(&mut a);
    println!("{:?}", a);
}