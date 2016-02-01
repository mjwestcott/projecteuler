"""
problem78.py

Let p(n) represent the number of different ways in which n coins can be
separated into piles. For example, five coins can be separated into piles
in exactly seven different ways, so p(5)=7.

         OOOOO
        OOOO   O
        OOO   OO
      OOO   O   O
      OO   OO   O
     OO   O   O   O
    O   O   O   O   O

Find the least value of n for which p(n) is divisible by one million.
"""
from itertools import count, cycle, takewhile
from toolset import memoize, take

def pentagonal(n):
    """Return nth pentagonal number e.g. 1, 5, 12, 22, 35, ..."""
    return n*(3*n - 1)/2

# https://en.wikipedia.org/wiki/Pentagonal_number_theorem
def generalised_pentagonals():
    x = 0
    while True:
        x += 1
        yield pentagonal(x); yield pentagonal(-x)

@memoize
def p(n):
    if n < 0:
        return 0
    elif n == 0:
        return 1
    else:
        # Generating pentagonals is repeated many times, should think about optimising this.
        pentagonals = list(takewhile(lambda x: x <= n, generalised_pentagonals()))
        terms = [p(n - x) for x in pentagonals]
        coefs = list(take(len(terms), cycle([1, 1, -1, -1])))
        return sum(a*b for (a, b) in zip(terms, coefs))

def problem78():
    return next(filter(lambda x: p(x) % 1000000 == 0, count(2)))
