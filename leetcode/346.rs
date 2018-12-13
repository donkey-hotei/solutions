/*
 * A simple implementation of circular queue (a.k.a: ring buffer).
 */

struct Queue {
    head: isize,
    tail: isize,
    data: Vec<isize>,
    capacity: usize,
}


impl Queue {
    fn with_capacity(capacity: usize) -> Self {
        let head = -1;
        let tail = -1;
        let data = Vec::with_capacity(capacity);

        Queue { 
            data,
            head,
            tail,
            capacity,
        }
    }

    fn is_full(&self) -> bool {
        (self.tail + 1) % self.capacity as isize == self.head
    }

    fn is_empty(&self) -> bool {
        self.head == -1
    }

    fn len(&self) -> usize {
        self.data.len()
    }

    fn enqueue(&mut self, element: isize) -> bool {
        if self.is_full() {
            return false;
        } else if self.is_empty() {
            self.head = 0;
        }

        self.tail = (self.tail + 1) % self.capacity as isize;
        self.data.insert(self.tail as usize, element);

        true
    }

    fn dequeue(&mut self) -> isize {
        if self.is_empty() {
            panic!("Attempted to dequeue from an empty queue!");
        }

        let element = self.data[self.head as usize];

        if self.head == self.tail {
            self.head = -1;
            self.tail = -1;
        }

        self.head = (self.head + 1) % self.capacity as isize;

        element
    }
}

struct MovingAverage {
    size: usize,
    window: Queue
}

impl MovingAverage {
    fn new(size: usize) -> Self {
        MovingAverage { size, window: Queue::with_capacity(size) }
    }

    fn next(&mut self, value: isize) -> f64 {
        if self.window.is_full() {
            self.window.dequeue();
        }

        self.window.enqueue(value);

        (self.window.data.iter().sum::<isize>() as f64) / self.window.len() as f64
    }
}

fn main() {
    let mut avg = MovingAverage::new(5);
    let stream = vec![1, 2, 3, 4, 5, 6, 1, 3, 5];

    for &value in stream.iter() {
        println!("{:?}", avg.next(value));
    }
}
