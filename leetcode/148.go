package main

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    return merge_sort(head)
}

func merge_sort(head *ListNode) *ListNode {
    // base case: list is empty or of length 1
    if head == nil || head.Next == nil {
        return head
    }

    // get middle node of the list, as initial pivot
    mid_node := get_middle(head)
    mid_succ := mid_node.Next

    // divide the list into two halves
    mid_node.Next = nil

    // and recursively sort
    left  := merge_sort(head)
    right := merge_sort(mid_succ)

    return sorted_merge(left, right)
}

func sorted_merge(a, b *ListNode) *ListNode {
    var result *ListNode

    if a == nil { return b }
    if b == nil { return a }

    if a.Val <= b.Val {
        result = a
        result.Next = sorted_merge(a.Next, b)
    } else {
        result = b
        result.Next = sorted_merge(a, b.Next)
    }

    return result
}

func get_middle(head *ListNode) *ListNode {
    if head == nil { return head }

    fast := head.Next
    slow := head

    for fast != nil {
        fast = fast.Next

        if fast != nil {
            slow = slow.Next
            fast = fast.Next
        }
    }

    return slow
}

