use std::collections::BinaryHeap;

fn find_kth_largest(v: Vec<i32>, k: i32) -> i32 {
    let mut heap = BinaryHeap::new();

    for element in v {
        heap.push(element);
    }

    for _ in 0..k - 1 {
        heap.pop();
    }

    *heap.peek().unwrap()
}

fn main() {
    let v = vec![3,2,1,5,6,4];
    let k = 2;

    assert_eq!(find_kth_largest(v, k), 5);
}