use std::cmp::Ordering;

#[derive(Debug, PartialEq, Eq)]
struct Interval {
    start: i32,
    end: i32,
}

impl Interval {
    fn new(start: i32, end: i32) -> Self {
        Interval { start, end }
    }
}

impl PartialOrd for Interval {
    fn partial_cmp(&self, other: &Interval) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Interval {
    fn cmp(&self, other: &Interval) -> Ordering {
        self.start.cmp(&other.start)
    }
}

fn can_attend_meetings(intervals: &mut Vec<Interval>) -> bool {
    if intervals.len() <= 1 { return false; }

    intervals.sort_by(|a, b| a.partial_cmp(b).unwrap());

    for i in 0..intervals.len() - 1 {
        if intervals[i].end > intervals[i + 1].start {
            return false;
        }
    }

    true
}

fn main() {
    let mut intervals = vec![
        Interval::new(5, 8),
        Interval::new(6, 8),
    ];
    assert!(can_attend_meetings(&mut intervals))
}