package main

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    if head.Next == nil || head.Next.Next == nil {
        return false
    }

    turtle := head
    rabbit := head.Next.Next

    for turtle != rabbit {
        if rabbit.Next == nil || rabbit.Next.Next == nil {
            return false
        }

        turtle = turtle.Next
        rabbit = rabbit.Next.Next
    }

    return false
}