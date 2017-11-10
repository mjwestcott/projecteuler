"""
problem64.py

https://projecteuler.net/problem=64

The first ten continued fraction representations of (irrational) square roots are:

    sqrt(2)=[1;(2)]          period=1
    sqrt(3)=[1;(1,2)]        period=2
    sqrt(5)=[2;(4)]          period=1
    sqrt(6)=[2;(2,4)]        period=2
    sqrt(7)=[2;(1,1,1,4)]    period=4
    sqrt(8)=[2;(1,4)]        period=2
    sqrt(10)=[3;(6)]         period=1
    sqrt(11)=[3;(3,6)]       period=2
    sqrt(12)=[3;(2,6)]       period=2
    sqrt(13)=[3;(1,1,1,1,6)] period=5

Exactly four continued fractions, for N <= 13, have an odd period. How many
continued fractions for N <= 10000 have an odd period?
"""
from math import floor, sqrt

from toolset import quantify


def continued_fraction_sqrt(S):
    # https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
    # Using variables S, m, d, a as in the URL above.
    m = 0
    d = 1
    a = floor(sqrt(S))
    seen = []
    while True:
        seen.append([m, d, a]) # The algorithm terminates when [m, d, a] repeats
        m = (d * a) - m
        d = (S - m**2) / d
        if d == 0: # S is a perfect square
            return [a]
        a = floor((floor(sqrt(S)) + m) / d)
        if [m, d, a] in seen:
            return [x[2] for x in seen] # The third element is the variable 'a' we want.

def problem64():
    continued_fractions = (continued_fraction_sqrt(i) for i in range(2, 10000+1))
    odd_period = lambda x: len(x) % 2 == 0 # The first element is not part of the period.
    return quantify(continued_fractions, pred=odd_period)

if __name__ == "__main__":
    print(problem64())
