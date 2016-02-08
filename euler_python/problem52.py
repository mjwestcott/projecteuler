"""
problem52.py

https://projecteuler.net/problem=52

It can be seen that the number, 125874, and its double, 251748, contain exactly
the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
contain the same digits.
"""
from itertools import count

def multiples(x):
    return [i*x for i in range(2, 6+1)]

def same_digits(x, y):
    return sorted(str(x)) == sorted(str(y))

def problem52():
    return next(x for x in count(1) if all(same_digits(x, y) for y in multiples(x)))
