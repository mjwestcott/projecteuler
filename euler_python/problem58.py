"""
problem58.py

Starting with 1 and spiralling anticlockwise in the following way, a square
spiral with side length 7 is formed.

    37 36 35 34 33 32 31
    38 17 16 15 14 13 30
    39 18  5  4  3 12 29
    40 19  6  1  2 11 28
    41 20  7  8  9 10 27
    42 21 22 23 24 25 26
    43 44 45 46 47 48 49

It is interesting to note that the odd squares lie along the bottom right
diagonal, but what is more interesting is that 8 out of the 13 numbers lying
along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.

If one complete new layer is wrapped around the spiral above, a square spiral
with side length 9 will be formed. If this process is continued, what is the
side length of the square spiral for which the ratio of primes along both
diagonals first falls below 10%?
"""
from itertools import count
from math import sqrt
from toolset import is_prime, quantify

def square_length(n):
    "Given the bottom right corner number, return the square length"
    return int(sqrt(n))

def corners(n):
    "Given the bottom right corner number, return the four corner numbers"
    # 49 --> [49, 43, 37, 31]
    x = square_length(n) - 1
    return [n, n-x, n-(2*x), n-(3*x)]

def problem58():
    # Yields all four corners from each new layer, starting at fifth layer.
    # next(all_corners) --> [81, 73, 65, 57], [121, 111, 101, 91], ...
    all_corners = (corners(x**2) for x in count(start=9, step=2))
    primes, total = 8, 13
    while True:
        cs = next(all_corners)
        primes += quantify(cs, pred=is_prime)
        total += 4
        if primes / total < 0.10:
            # cs[0] is the bottom right corner number
            return square_length(cs[0])
