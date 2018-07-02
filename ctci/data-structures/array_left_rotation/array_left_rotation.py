#!/usr/bin/python

def left_rotate_by_one(a, n):
	tmp = a[0]
	for i in xrange(0, n - 1):
		a[i] = a[i+1]
	a[-1] = tmp
	return a


def left_rotate_list(a, n, k):
	for i in xrange(k):
		a = left_rotate_by_one(a, n)
	return a

if __name__ == '__main__':
	import sys
	n, k = map(int, sys.stdin.readline().split())
	a = map(int, sys.stdin.readline().split())
	r = left_rotate_list(a, n, k)
	for i in r:
		print i,
