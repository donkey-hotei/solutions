class Stack:
    def __init__(self):
        self.elements = []

    def push(self, value):
        self.elements.append(value)

    def pop(self):
        return self.elements.pop()

    def top(self):
        return self.elements[-1]

    def is_empty(self):
        return len(self.elements) == 0


class Node:
    def __init__(self, value):
        self.value = value
        self.left, self.right = None, None

    def insert_left(self, value):
        if self.left is not None:
            return self.left.insert(value)
        else:
            self.left = Node(value)
            return self.left

    def insert_right(self, value):
        if self.right is not None:
            return self.right.insert(value)
        else:
            self.right = Node(value)
            return self.right

    def insert(self, value):
        if value <= self.value:
            return self.insert_left(value)
        else:
            return self.insert_right(value)

    def kth_smallest(self, k):
        array = []
        stack = Stack()

        current = self

        while current is not None or not stack.is_empty() and len(array) < k:
            while current is not None:
                stack.push(current)
                current = current.left

            current = stack.pop()
            array.append(current.value)
            print(array)
            current = current.right

        return array[k - 1]


def print_inorder(root):
    if root is None:
        return
    print_inorder(root.left)
    print("{} ".format(root.value)),
    print_inorder(root.right)


def main():
    root = Node(5)
    root.insert(3)
    root.insert(6)
    root.insert(2)
    root.insert(4)
    root.insert(1)
    print_inorder(root)
    print("kth smallest: {}".format(root.kth_smallest(1)))


if __name__ == "__main__":
    main()
