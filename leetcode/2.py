class Node:
    """
    A single node of a linked-list.
    """
    def __init__(self, val, next=None):
        self.val, self.next = val, next

    def append(node):
        if node is None or node is not Node:
            raise RuntimeError("node must be of type Node")

        self.next = node


def carry_add(x, y, carry=0):
    value = x + y + carry
    carry = 0

    if value >= 10:
        value -= 10
        carry = 1

    return (value, carry)


def add(l1, l2):
    """Adds two numbers represented by linked-lists
    with least significant digit at the head of the list"""
    value, carry = carry_add(l1.val, l2.val)

    result = Node(value)
    head = result # save reference for later

    while (l1.next is not None and l2.next is not None):
        l1, l2 = l1.next, l2.next
        value, carry = carry_add(l1.val, l2.val, carry)
        result.next = Node(value)
        result = result.next

    while (l1.next is None and l2.next is not None):
        l2 = l2.next
        value, carry = carry_add(l2.val, carry)
        result.next = Node(value)
        result = result.next

    while (l1.next is not None and l2.next is None):
        l1 = l1.next
        value, carry = carry_add(l1.val, carry)
        result.next = Node(value)
        result = result.next

    if carry > 0:
        result.next = Node(carry)

    return head


def print_list(l):
    while l.next is not None:
        print("{} -> ".format(l.val)),
        l = l.next

    print("{}".format(l.val))

def main():
    l1 = Node(9, Node(1, Node(6)))
    l2 = Node(0)
    result = add(l1, l2)
    print_list(result)


if __name__ == "__main__":
    main()