struct MinStack(Vec<i32>);


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {

    /** initialize your data structure here. */
    fn new() -> Self {
        MinStack(Vec::new())
    }

    fn push(&mut self, x: i32) {
        self.0.push(x)
    }

    fn pop(&mut self) {
        self.0.pop();
    }

    fn top(&self) -> i32 {
        self.0[self.0.len() - 1]
    }

    fn get_min(&self) -> i32 {
        if let Some(&min) = self.0.iter().min() {
            min
        } else {
            panic!("Stack is empty!")
        }
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * let obj = MinStack::new();
 * obj.push(x);
 * obj.pop();
 * let ret_3: i32 = obj.top();
 * let ret_4: i32 = obj.get_min();
 */

fn main() {
    let mut obj = MinStack::new();

    obj.push(-2);
    obj.push(0);
    obj.push(-3);
    dbg!(obj.get_min());
    obj.pop();
    dbg!(obj.top());
    dbg!(obj.get_min());
}