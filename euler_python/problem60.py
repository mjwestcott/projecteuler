"""
problem60.py

The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes
and concatenating them in any order the result will always be prime. For
example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four
primes, 792, represents the lowest sum for a set of four primes with this
property.

Find the lowest sum for a set of five primes for which any two primes
concatenate to produce another prime.
"""
from itertools import takewhile
from toolset import memoize, is_prime, get_primes

@memoize
def concats_to_prime(x, y):
    "Tests whether concatenating x and y in either order makes a prime"
    check = lambda x, y: is_prime(int(str(x) + str(y)))
    return check(x, y) and check(y, x)

def all_concat_to_prime(candidates):
    return all(concats_to_prime(x, y)
               for x in candidates
               for y in candidates
               if x != y)

# It's not clear how many prime numbers to search through.
# Running next(n for n in count(start=0, step=1000) if problem60(n))
# suggests 9000.
def problem60(limit=9000):
    primes = list(takewhile(lambda x: x < limit, get_primes()))

    # Use depth-first tree search.
    frontier = [[p] for p in reversed(primes)] # check smaller primes first
    while frontier:
        node = frontier.pop()
        if len(node) == 5:
            return sum(node)
        for x in primes:
            child = node + [x]
            if x > max(node) and all_concat_to_prime(child):
                frontier.append(child)
