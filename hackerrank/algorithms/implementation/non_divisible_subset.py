#!/usr/bin/python
from math import ceil

if __name__ == '__main__':
	import sys
	n, k = map(int, sys.stdin.readline().split())
	s = set(map(int, sys.stdin.readline().split()))
	r = { n % k for n in s }
	S = set() 
	for k in range(1, math.ceil(k / 2) - 1):
		# if count(R, k) >= count(R, K - k)
                #         add all n_i to S, s.t. r_i == k
                # else:
                #         add all n_i to S, s.t. r_i == K - k
		pass
	# add only one n_i to S, s.t. r_i == 0, if exists
	# if K is even:
	#         add only one n_i to S s.t. r_i == K / 2, if exists
	# output S
	
