struct NumArray {
    data: Vec<i32>
}


/** You can modify the type of `self` for your need. */
impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        NumArray { data: nums }
    }

    fn sum_range(&self, i: i32, j: i32) -> i32 {
        self.data[i as usize..=j as usize].iter().sum()
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * let obj = NumArray::new(nums);
 * let ret_1: i32 = obj.sum_range(i, j);
 */
