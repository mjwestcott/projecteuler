"""
problem76.py

https://projecteuler.net/problem=76

It is possible to write five as a sum in exactly six different ways:

    4 + 1
    3 + 2
    3 + 1 + 1
    2 + 2 + 1
    2 + 1 + 1 + 1
    1 + 1 + 1 + 1 + 1

How many different ways can one hundred be written as a sum of at least two positive integers?
"""
from toolset import memoize

@memoize
def num_partitions(n, k):
    """Return the number of partitions of n, using positive integers <= k"""
    # I found this much easier to understand as 'change-giving' (as in problem number 31).
    # Solving num_partitions(n=100, k=99) means solving the number of ways to give change
    # to 100 using values in the set {1, 2, 3, ..., 99}. This can be broken down into
    # sub-problems, as the answer is the sum of the ways to give change
    # to 99, i.e. n-1, since we can start by using 1
    #    98, i.e. n-2, since we can start by using 2
    #    ...
    #    1, i.e. n-99, since we can start by using 99
    # But simply recursively summing all those ways to give change will
    # over-count many solutions. For instance, 5 = 3 + 1 + 1 is the same
    # as 5 = 1 + 1 + 3. So we need to determine a canonical way to give change.
    # This can be achieved by specifying that having used a coin of value x as
    # the first step, we can only use coins of value <= x from then on.
    # So, the solution to 99, i.e. n-1, can use only {1}
    #                     98, i.e. n-2, can use values in the set {1, 2}
    #                     97, i.e. n-3, can use values in the set {1, 2, 3}, etc.
    # This is how we arrive at sum(num_partitions(n-x, x) for x in range(1, k+1)) below.
    if n < 0:
        # This will occur after an attempt to give change for n, with a coin
        # greater than n, and indicates the failure of change-giving.
        return 0
    elif n == 0:
        # This will occur after an attempt to give change for n, with a coin
        # of value exactly n, and indicates the change-giving was successful.
        return 1
    else:
        # For all possible coin-values, x, find the ways to give change to
        # (n-x) using coins <= x.
        return sum(num_partitions(n-x, x) for x in range(1, k+1))

def problem76():
    return num_partitions(100, 99)
