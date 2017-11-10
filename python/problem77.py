"""
problem77.py

https://projecteuler.net/problem=77

It is possible to write ten as the sum of primes in exactly five different ways:

    7 + 3
    5 + 5
    5 + 3 + 2
    3 + 3 + 2 + 2
    2 + 2 + 2 + 2 + 2

What is the first value which can be written as the sum of primes in over five
thousand different ways?
"""
from itertools import count, takewhile

from toolset import get_primes, memoize_mutable


@memoize_mutable
def num_partitions(n, primes):
    # Using a slightly different algorithm than problem 76.
    # This one is adapted from SICP: https://mitpress.mit.edu/sicp/full-text/book/book-Z-H-11.html
    # See the section entitled 'Example: Counting change'. Their logic is
    # more intuitive than that which I presented in the previous problem.
    if n < 0:
        return 0
    elif n == 0:
        return 1
    elif primes == []:
        return 0
    else:
        return num_partitions(n, primes[1:]) + num_partitions(n - primes[0], primes)

def problem77():
    primes = list(takewhile(lambda x: x < 100, get_primes()))
    return next(filter(lambda x: num_partitions(x, primes) > 5000, count(2)))

if __name__ == "__main__":
    print(problem77())
