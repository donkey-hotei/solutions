#!/usr/bin/python


if __name__ == '__main__':
    import sys
    n = int(sys.stdin.readline())
    a = map(int, sys.stdin.readline().split())
    q = int(sys.stdin.readline())
    x = map(int, sys.stdin.readline().split())
    
    for i in x:
        a = map(lambda x: i + x, a)
        r = sum([abs(i) for i in a])
        print(r)

