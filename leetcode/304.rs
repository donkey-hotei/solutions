struct NumMatrix {
    matrix: Vec<Vec<i32>>,
}


impl NumMatrix {
    fn new(matrix: Vec<Vec<i32>>) -> Self {
        NumMatrix { matrix }
    }

    fn sum_region(&self, row1: i32, col1: i32, row2: i32, col2: i32) -> i32 {
        (row1..=row2)
            .flat_map(|x| (col1..=col2).map(move |y| (x, y)))
            .fold(0, |acc, (x, y)| {
                acc + self.matrix[x as usize][y as usize]
            })
    }
}