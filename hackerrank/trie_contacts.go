package main

/*
 * Represents a Node within a Trie
 */
type Node struct {
    label string
    data  string
    children map[rune]*Node
}

/*
 * A simple Trie starting at some *Node head.
 */
type Trie struct {
     head *Node
}

/*
 * Queue (FIFO) based on a circular list.
 */
type Queue struct {
    nodes []*Node
    head  int
    tail  int
    count int
    size  int
}

func (q *Queue) Push() {
    // add element to queue
}


func (q *Queue) Pop() {
    // pop element from queue
}

/*
 * Finds all words in the trie  which whose prefix match.
 */
func (t *Trie) find(prefix string) []string {
    words := make([]string, 0)

    curr := t.head;
    for c, _ := range prefix {
        curr, ok := t.head.children[c]; if !ok {
            return words
        }
    }

    // use queue to perform a depth first search
    queue = make([]*Node, 0)

    for !queue.isEmpty {
        curr := queue.Pop()

    }
}



func main() {

}