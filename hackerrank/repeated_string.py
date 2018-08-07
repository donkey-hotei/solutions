#!/usr/bin/python

if __name__ == '__main__':
	import sys
	s = sys.stdin.readline()
	k = len(s)
	n = int(sys.stdin.readline())
	c = 0

	for i in s:
		if i == 'a':
			c += 1

	ans = c * (n // k)

	for i in s[:k % n + 1]:
		if i == 'a':
			ans += 1	

	print(ans)
	
