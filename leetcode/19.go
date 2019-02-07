/* leetcode #19: Remove Nth Node from Singly Linked List */
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func length(head *ListNode) int {
    count := 0

    for head != nil {
        count += 1
        head = head.Next
    }

    return count
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    listlen := length(head)
    listlen -= n

    dummy := &ListNode { Val: 0, Next: head }
    cur := dummy

    for listlen > 0 {
        listlen -= 1
        cur = cur.Next
    }

    cur.Next = cur.Next.Next

    return dummy.Next
}
