struct NumMatrix {
    cache:  Vec<Vec<i32>>,
}

impl NumMatrix {
    fn new(matrix: Vec<Vec<i32>>) -> Self {
        let (m, n) = (matrix.len(), matrix[0].len());

        if m == 0 || n == 0 { return NumMatrix { cache: vec![] } }

        let mut cache = vec![vec![0; n + 1]; m];
        for x in (0..m) {
            for y in (0..n) {
                cache[x][y + 1] = cache[x][y] + matrix[x][y];
            }
        }

        println!("{:?}", cache);

        NumMatrix { cache }
    }

    fn sum_region(&self, row1: i32, col1: i32, row2: i32, col2: i32) -> i32 {
        if self.cache.len() == 0 { return 0; }

        (row1..=row2).fold(0, |acc, row| {
            acc +
                self.cache[row as usize][col2 as usize + 1] -
                self.cache[row as usize][col1 as usize]
        })
    }
}

fn main() {
    let matrix = vec![
        vec![3, 0, 1, 4, 2],
        vec![5, 6, 3, 2, 1],
        vec![1, 2, 0, 1, 5],
        vec![4, 1, 0, 1, 7],
        vec![1, 0, 3, 0, 5]
    ];

    // let empty_matrix = vec![vec![]];

    let obj = NumMatrix::new(matrix);

    println!("{:?}", obj.sum_region(2, 1, 4, 3));
}