fn can_partition(a: Vec<i32>) -> bool {
    let sum: i32 = a.iter().sum();

    if sum % 2 == 1 { return false; }

    let mut possible_partition =
        vec![false; (sum as usize / 2) + 1];

    possible_partition[0] = true;

    for &e in a.iter() {
        for i in ((e as usize)..=possible_partition.len() - 1).rev() {
            if possible_partition[i - e as usize] {
                possible_partition[i] = true;
            }
        }
    }

    possible_partition[possible_partition.len() - 1]
}

fn main() {
    let a = vec![1, 2, 1, 3, 5];

    println!("{:?}", can_partition(a));
}