"""
Detect a cycle in a linked list. Note that the head pointer may be 'None' if the list is empty.

A Node is defined as: 
"""

class Node(object):
    def __init__(self, data = None, next_node = None):
        self.data = data
        self.next = next_node


def has_cycle(head):
    assert(type(head) == Node)

    tail = head.next

    # > 2 nodes
    if tail.next is not None:
        tail = tail.next

    while head and tail:
        if head is tail:
            return True

        t = t.next
        if h.next:
            h = h.next.next

    return False


if __name__ == '__main__':
    import sys

    ls = Node()
    for i in range(3):
        l = ls.next
        if l is None:
            l = Node(data=i)
            ls.next = l

    if has_cycle(ls):
        print "Cycle detected."
    else:
        print "No cycle detected."

