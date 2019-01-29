fn is_saddle_point(m: &[Vec<u64>], x: usize, y: usize) -> bool {
    for row in 0..m.len() {
        if m[row][y] < m[x][y] {
            return false
        }
    }

    for col in 0..m[0].len() {
        if m[x][col] > m[x][y] {
            return false
        }
    }

    true
}

pub fn find_saddle_points(matrix: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let mut points = Vec::new();

    for x in 0..matrix.len() {
        for y in 0..matrix[0].len() {
            if is_saddle_point(&matrix, x, y) {
                points.push((x, y));
            }
        }
    }

    points
}
