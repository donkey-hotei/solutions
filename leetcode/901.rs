struct Node<T> {
    data: T,
    next: Option<Box<Node<T>>>,
}

struct List<T> {
    head: Option<Box<Node<T>>>,
}

impl<T> List<T> {
    fn new() -> Self {
        List { head: None }
    }

    fn push(&mut self, elem: T) {
        self.head = Some(Box::new(Node { data: elem, next: self.head.take() }));
    }

    fn pop(&mut self) -> Option<T> {
        self.head.take().map(|mut head| {
            self.head = head.next.take();
            head.data
        })
    }

    fn peek(&self) -> Option<&T> {
        if let Some(ref node) = self.head {
            Some(&node.data)
        } else {
            None
        }
    }
}

struct StockSpanner {
    stack: List<(i32, i32)>
}

impl StockSpanner {
    fn new() -> Self {
        StockSpanner { stack: List::new() }
    }

    fn next(&mut self, price: i32) -> i32 {
        let mut weight = 1;

        while let Some(&(p, w)) = self.stack.peek() {
            if price >= p {
                self.stack.pop();
                weight += w;
            } else {
                break;
            }
        }
        self.stack.push((price, weight));

        weight
    }
}

fn main() {
    let prices = vec![100, 80, 60, 70, 60, 75, 85];
    let mut spanner = StockSpanner::new();

    for price in prices {
        println!("{:?}", spanner.next(price))
    }
}