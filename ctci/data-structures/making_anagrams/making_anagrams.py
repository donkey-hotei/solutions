#!/usr/bin/python
from collections import Counter

def number_of_deletions_needed(s1, s2):
    d1 = Counter(s1)
    d2 = Counter(s2)
    c = d1 - d2
    d = d2 - d1
    e = c + d
    deletions = len(list(e.elements()))
    return deletions

if __name__ == '__main__':
    s1 = 'canmligyxyvym'
    s2 = 'jxwtrhvujlmrpdoqbisbwhmgpmeoke'
    print number_of_deletions_needed(s1, s2)

