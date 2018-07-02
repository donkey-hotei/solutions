#!/usr/bin/python
from collections import Counter



if __name__ == '__main__':
    import sys
    m, n = map(int, sys.stdin.readline().split())
    newspaper = Counter(sys.stdin.readline().strip().split(' '))
    ransom_note = Counter(sys.stdin.readline().strip().split(' '))
    result = all([ransom_note[word] <= newspaper[word] for word in ransom_note.keys()])
    if result:
        print "Yes"
    else:
        print "No"

