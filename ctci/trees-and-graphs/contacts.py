import collections

class Node(object):
    def __init__(self, label=None, data=None):
        self.label, self.data = label, data
        self.children = collections.defaultdict()

    def add_child(self, label, data=None):
        self.children[label] = Node(label, data)

    def __getitem__(self, key):
        return self.children[key]

    def __str__(self):
        return "Node({0}, {1})".format(self.label, self.data)


class Trie(object):
    def __init__(self, *args):
        self.head = Node()

    def __getitem__(self, key):
        return self.head.children[key]

    def find(self, word):
        """ Counts all words in the trie which
            have at least word as their shared prefix.
        """

        curr = self.head
        words = []
        # Do we at least contain the whole word?
        for letter in word:
            if letter in curr.children:
                curr = curr.children[letter]
            else:
                return words

        queue = [curr]

        while len(queue):
            curr = queue.pop()

            if "_end" in curr.children:
                words.append(curr.data)

            queue = [node
                     for _, node in
                     curr.children.iteritems()] + queue

        return words

    def insert(self, word):
        """ Inserts a word letter-by-letter into
            the trie.
        """
        curr = self.head
        last_letter = 0
        # Move along the trie for as many nodes that match
        # letters we have in the word.
        for letter in word:
            if letter in curr.children:
                curr = curr.children[letter]
                last_letter += 1
            else:
                break

        # rest_of_world always <= word
        rest_of_word = word[last_letter:]

        if len(rest_of_word):
            # then, the rest are added as new children
            for i, letter in enumerate(rest_of_word):
                curr.add_child(letter, word[:i + 1])
                curr = curr[letter]
            # mark the last node as a terminating node
            curr.data = word
            curr.add_child("_end")


if __name__ == "__main__":
    import sys

    t = Trie()
    n = map(int, sys.stdin.readline().split())[0]

    for _ in range(n):
        operation, word = sys.stdin.readline().split()

        if operation == "add":
            t.insert(word)
        elif operation == "find":
            print len(t.find(word))
        else:
            continue

