enum Solution { }

impl Solution {
    fn find_the_difference(s: String, t: String) -> char {
        let mut xor = 0;

        for &byte in s.as_bytes() {
            xor ^= byte
        }
        for &byte in t.as_bytes() {
            xor ^= byte
        }

        xor as char
    }
}

fn main() {
    println!("{:?}", Solution::find_the_difference("abcd".to_string(), "abcde".to_string()))
}