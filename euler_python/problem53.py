"""
problem53.py

https://projecteuler.net/problem=53

How many, not necessarily distinct, values of nCr, for 1 ≤ n ≤ 100, are greater
than one-million?
"""
from math import factorial

def num_combinations(n, r):
    return factorial(n) / (factorial(r) * factorial(n - r))

def problem53():
    return sum(1 for n in range(1, 100+1)
                 for r in range(1, n+1)
                 if num_combinations(n, r) > 1e6)
