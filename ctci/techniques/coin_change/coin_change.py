#!/usr/bin/env python

#!/bin/python3

import sys

def make_change(coins, n):
    pass

if __name__ == "__main__":
    n,m = input().strip().split(' ')
    n,m = [int(n),int(m)]
    coins = [int(coins_temp) for coins_temp in input().strip().split(' ')]
    print(make_change(coins, n))

