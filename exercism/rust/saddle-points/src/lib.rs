struct MatrixIterator {
    matrix: Vec<Vec<u64>>
}

impl MatrixIterator {
    fn new(matrix: Vec<Vec<u64>>) -> Self {
        MatrixIterator { matrix }
    }

    fn enumerate(&self) -> impl Iterator<Item = (usize, usize)> + '_ {
        self.matrix
            .iter()
            .enumerate()
            .flat_map(|(x, v)| {
                v.iter()
                 .enumerate()
                 .map(move |(y, _)| (x, y))
            })
    }
}

fn is_saddle_point(m: &[Vec<u64>], x: usize, y: usize) -> bool {
    if (0..m.len()).any(|row| m[row][y] < m[x][y]) {
        false
    } else if (0..m[0].len()).any(|col| m[x][col] > m[x][y]) {
        false
    } else {
        true
    }
}

pub fn find_saddle_points(matrix: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let matrix_iterator = MatrixIterator::new(matrix.into());

    matrix_iterator
        .enumerate()
        .fold(Vec::new(), |mut points, (x, y)| {
            match is_saddle_point(matrix, x, y) {
                false => points,
                true => {
                    points.push((x, y));
                    points
                }
            }
        })
}
